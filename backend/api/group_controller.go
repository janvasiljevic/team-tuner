package api

import (
	"encoding/csv"
	"fmt"
	"jv/team-tone-tuner/dto/in"
	"jv/team-tone-tuner/dto/out"
	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/model/course"
	"jv/team-tone-tuner/model/group"
	"jv/team-tone-tuner/model/grouprun"
	"jv/team-tone-tuner/model/user"
	"jv/team-tone-tuner/teamform"
	"jv/team-tone-tuner/utils"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

//	@Summary		Get groups
//	@Description	Get groups
//	@ID				get-groups-by-group-run
//	@Tags			groups
//	@Param			request	query		in.GetGroupsQuery	true	"Get groups query"
//	@Success		200		{array}		out.GroupOut		"Groups filtered by query"
//	@Failure		400		{object}	out.Error			"Bad request"
//	@Failure		401		{object}	out.Error			"Unauthorized"
//	@Failure		403		{object}	out.Error			"Forbidden"
//	@Failure		500		{object}	out.Error			"Internal server error"
//	@Router			/group [get]
func (api Api) GetGroups(c echo.Context) error {
	ctx := c.Request().Context()

	var dto in.GetGroupsQuery

	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequest())
	}

	if err := c.Validate(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, out.NewValidatorError(err))
	}

	query := api.db.Group.Query().WithStudents().WithCourse()

	log.Info().Msg(fmt.Sprintf("dto: %+v", dto))

	if dto.GroupRun != nil {
		query = query.Where(
			group.HasGroupRunWith(grouprun.IDEQ(*dto.GroupRun)),
		)
	}

	groups, err := query.All(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to get groups")
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	groupsOut := utils.Map(groups, out.ConvertModelToGroupOut)

	return c.JSON(http.StatusOK, groupsOut)
}

//	@Summary		Get group by id: their students and stats
//	@Description	Get group by id: their students and stats (bfi scores)
//	@ID				get-group-by-id
//	@Tags			groups
//	@Param			id	path		string	true	"Group run id"
//	@Success		200	{object}	out.DetailedGroupOut
//	@Router			/group/{id} [get]
func (api Api) GetGroup(c echo.Context) error {
	ctx := c.Request().Context()
	rawId := c.Param("id")

	gId, err := uuid.Parse(rawId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequestWithMessage("Invalid group id"))
	}

	group, err := api.db.Group.
		Query().
		Where(group.IDEQ(gId)).
		WithStudents(func(q *model.UserQuery) {
			q.WithBfiReport()
		}).
		WithCourse().
		Only(ctx)

	if err != nil {
		return c.JSON(http.StatusNotFound, out.NotFound())
	}

	return c.JSON(http.StatusOK, out.NewDetailedGroupOut(group))
}

//	@Summary		Create group
//	@Description	Create groups for a course. Runs an algorithm to assign students to groups. After the algorithm finished the following are created: A group run and groups with students asiigned to them
//	@ID				create-groups
//	@Tags			groups
//	@Param			request	body	in.PostFormGroupsIn	true	"Create groups request"
//	@Success		200		"OK"
//	@Failure		400		{object}	out.Error	"Bad request"
//	@Failure		401		{object}	out.Error	"Unauthorized"
//	@Failure		403		{object}	out.Error	"Forbidden"
//	@Failure		500		{object}	out.Error	"Internal server error"
//	@Router			/group/form [post]
func (api Api) PostFormGroups(c echo.Context) error {

	var dto in.PostFormGroupsIn

	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequest())
	}

	if err := c.Validate(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, out.NewValidatorError(err))
	}

	ctx := c.Request().Context()
	adminId := c.Get("userId").(uuid.UUID)

	log.Info().Str("adminId", adminId.String()).Str("courseId", dto.CourseId.String()).
		Int("iterations", dto.Iterations).
		Float64("temperature", dto.Temperature).
		Float64("weight satisfaction", dto.WeightSatisfaction).
		Float64("weight neuroticism", dto.WeightNeuroticism).
		Float64("weight extraversion", dto.WeightExtraversion).
		Float64("weight conscientiousness", dto.WeightConscientiousness).
		Msg("Creating groups")

	students, err := api.db.User.
		Query().
		Where(
			user.HasCoursesWith(course.IDEQ(*dto.CourseId)),
			user.RoleEQ(user.RoleStudent),
		).
		WithBfiReport().
		All(ctx)

	if err != nil {
		log.Error().Str("adminId", adminId.String()).Str("courseId", dto.CourseId.String()).
			Err(err).Msg("Failed to get students from DB")

		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	startTime := time.Now()

	res, err := teamform.GenerateTeams(students,
		dto.GroupSizes,
		dto.Iterations,
		dto.Temperature,
		dto.WeightSatisfaction,
		dto.WeightNeuroticism,
		dto.WeightExtraversion,
		dto.WeightConscientiousness,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to generate groups - algorithm error")
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	// create a group run
	groupRun, err := api.db.GroupRun.Create().
		SetCourseID(*dto.CourseId).
		SetBestScoreHistory(res.BestScoreHistory).
		SetCandidateScoreHistory(res.BestScoreHistory).
		SetGroupSize(dto.GroupSizes).
		SetCourseID(*dto.CourseId).
		SetCreatedByID(adminId).
		Save(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to save group run to DB")
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	log.Info().Str("group run id", groupRun.ID.String()).
		Int("sa iterations", len(res.BestScoreHistory)).
		Float64("sa time elapsed", time.Since(startTime).Seconds()).
		Msg("Saved group run to DB")

	createGroupsBulk := make([]*model.GroupCreate, len(res.Groups))

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i, group := range res.Groups {

		studentIds := utils.Map(group.Students, func(s teamform.Student) uuid.UUID {
			return s.Id
		})

		createGroupsBulk[i] = api.db.Group.Create().
			SetGroupRun(groupRun).
			AddStudentIDs(studentIds...).
			SetName(utils.GenerateRandomDockerLikeName(rnd)).
			SetCourseID(*dto.CourseId)
	}

	_, err = api.db.Group.CreateBulk(createGroupsBulk...).Save(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to save groups to DB")
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	log.Info().Str("group run id", groupRun.ID.String()).Int("groups len", len(createGroupsBulk)).Msg("Saved groups to DB")

	return c.NoContent(http.StatusOK)
}

//	@Summary		Download groups as CSV
//	@Description	Download groups as CSV
//	@ID				download-groups-csv
//	@Tags			groups
//	@Param			groupRunId	path	string	true	"Group run id"
//	@Success		200			"OK"
//	@Failure		400			{object}	out.Error	"Bad request"
//	@Failure		401			{object}	out.Error	"Unauthorized"
//	@Failure		403			{object}	out.Error	"Forbidden"
//	@Failure		404			{object}	out.Error	"Not found"
//	@Failure		500			{object}	out.Error	"Internal server error"
//	@Router			/group/download/{groupRunId} [get]
func (api Api) GetDownloadGroupsCSV(c echo.Context) error {

	param := c.Param("groupRunId")

	id, err := uuid.Parse(param)

	if err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequestUUID())
	}

	ctx := c.Request().Context()

	groups, err := api.db.GroupRun.Query().Where(grouprun.IDEQ(id)).QueryGroups().WithStudents().All(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to get groups from DB")
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	// Create a new CSV writer that writes to a string builder
	sb := &strings.Builder{}
	writer := csv.NewWriter(sb)

	err = writer.Write([]string{"groupId", "groupName", "studentGithub"})

	if err != nil {
		log.Error().Err(err).Msg("Failed to write headers to CSV")
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	for _, group := range groups {
		for _, student := range group.Edges.Students {

			err := writer.Write([]string{
				group.ID.String(),
				group.Name,
				student.GithubUsername,
			})

			if err != nil {
				log.Error().Err(err).Str("group id", group.ID.String()).Str("student id", student.ID.String()).Msg("Failed to write to CSV")
				return c.JSON(http.StatusInternalServerError, out.NewInternal())
			}
		}
	}

	writer.Flush()

	// Set the Content-Type and Content-Disposition headers to indicate a CSV file
	c.Response().Header().Set(echo.HeaderContentType, "text/csv")
	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=groups-%s.csv", id.String()))

	// Write the CSV data to the response
	return c.String(http.StatusOK, sb.String())
}
