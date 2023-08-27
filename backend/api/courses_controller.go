package api

import (
	"fmt"
	"jv/team-tone-tuner/dto/in"
	"jv/team-tone-tuner/dto/out"
	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/model/bfireport"
	"jv/team-tone-tuner/model/course"
	"jv/team-tone-tuner/model/grouprun"
	"jv/team-tone-tuner/model/user"
	"jv/team-tone-tuner/utils"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

//	@Summary		Get courses
//	@Description	Get courses
//	@Tags			courses
//	@Success		200	{array}	out.CourseOut
//	@Router			/course [get]
func (api Api) GetCourses(c echo.Context) error {

	ctx := c.Request().Context()

	courses, err := api.db.Course.Query().All(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to get courses")
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	coursesOut := utils.Map(courses, out.ConvertModelToCourseOut)

	return c.JSON(http.StatusOK, coursesOut)
}

//	@Summary		Create course
//	@Description	Create course
//	@Tags			courses
//	@Param			body	body		in.CreateCourseBody	true	"Create course body"
//	@Success		201		{object}	out.CreatedCourseOut
//	@Router			/course [post]
func (api Api) PostCourse(c echo.Context) error {
	ctx := c.Request().Context()

	var postData in.CreateCourseBody

	if err := c.Bind(&postData); err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequest())
	}

	if err := c.Validate(&postData); err != nil {
		return c.JSON(http.StatusBadRequest, out.NewValidatorError(err))
	}

	course, err := api.db.Course.Create().
		SetName(postData.Name).
		SetCode(postData.Code).
		SetColour(postData.Colour).
		Save(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	return c.JSON(http.StatusCreated, out.ConvertModelToCreatedCourseOut(course))
}

//	@Summary		Get a big five box plot for a course
//	@Description	Get a big five box plot for a course
//	@Tags			courses
//	@Param			courseId	path		string	true	"Course ID"
//	@Success		200			{object}	out.BigFiveBoxPlot
//	@Router			/course/{courseId}/stats/bf-box-plot [get]
func (api Api) GetBigFiveBoxPlot(c echo.Context) error {
	ctx := c.Request().Context()

	courseId := c.Param("courseId")

	uuid, err := uuid.Parse(courseId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequest())
	}

	reports, err := api.db.Course.Query().Where(
		course.ID(uuid),
	).QueryStudents().QueryBfiReport().All(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to get reports")
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	boxPlot := out.NewBigFiveBoxPlotFromReports(reports)

	return c.JSON(http.StatusOK, boxPlot)
}

//	@Summary		Get a daily number of solved questioners for a course
//	@Description	Get a daily number of solved questioners for a course
//	@Tags			courses
//	@Param			courseId	path	string						true	"Course ID"
//	@Param			request		query	in.GetDailyActivityRequest	true	"Get daily activity request"
//	@Router			/course/{courseId}/stats/daily-activity [get]
//	@Success		200	{object}	out.DailyActivity
func (api Api) GetDailyActivity(c echo.Context) error {

	ctx := c.Request().Context()

	var query in.GetDailyActivityRequest

	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequest())
	}

	if err := c.Validate(&query); err != nil {
		return c.JSON(http.StatusBadRequest, out.NewValidatorError(err))
	}

	rawParam := c.Param("courseId")
	uuid, err := uuid.Parse(rawParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequestWithMessage("Invalid course id"))
	}

	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()

	if query.StartTime != nil {
		startTime = *query.StartTime
	}

	if query.EndTime != nil {
		endTime = *query.EndTime
	}

	if startTime.After(endTime) {
		return c.JSON(http.StatusBadRequest, out.BadRequestWithMessage("Start time cannot be after end time"))
	}

	var returnedRows []out.RowReturn

	api.db.Course.Query().Where(course.ID(uuid)).
		QueryStudents().
		QueryBfiReport().
		Where(bfireport.CreatedAtGTE(startTime), bfireport.CreatedAtLTE(endTime)).
		Aggregate(func(s *sql.Selector) string {
			as := "day"
			grouper := fmt.Sprintf("date_trunc('day', %s)", bfireport.FieldCreatedAt)
			s.GroupBy(grouper)
			return sql.As(grouper, as)
		}).
		Aggregate(model.Count()).
		ScanX(ctx, &returnedRows)

	out := out.NewDailyActivityFromRows(returnedRows, startTime, endTime)

	return c.JSON(http.StatusOK, out)
}

//	@Summary		Get a questioner completion stats for a course
//	@Description	Get a questioner completion stats for a course
//	@Tags			courses
//	@Param			courseId	path		string	true	"Course ID"
//	@Success		200			{object}	out.QuestionerCompletionStats
//	@Router			/course/{courseId}/stats/questioner-stats [get]
func (api Api) GetQuestionerCompletionStats(c echo.Context) error {
	ctx := c.Request().Context()
	raw := c.Param("courseId")

	uuid, err := uuid.Parse(raw)

	if err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequest())
	}

	numOfStudents := api.db.Course.Query().Where(course.ID(uuid)).
		QueryStudents().Where(user.RoleEQ(user.RoleStudent)).CountX(ctx)

	numOfFinished := api.db.Course.Query().Where(course.ID(uuid)).
		QueryStudents().Where(user.RoleEQ(user.RoleStudent)).QueryBfiReport().CountX(ctx)

	return c.JSON(http.StatusOK, out.NewQuestionerCompletionStatsFromRows(numOfStudents, numOfFinished))
}

//	@Summary		Get courses group runs
//	@Description	Get courses group runs
//	@ID				get-courses-group-runs
//	@Tags			courses
//	@Param			courseId	path		string	true	"Course ID"
//	@Success		200			{object}	[]out.GroupRunOut
//	@Router			/course/{courseId}/group-runs [get]
func (api Api) GetGroupRuns(c echo.Context) error {
	ctx := c.Request().Context()
	raw := c.Param("courseId")

	uuid, err := uuid.Parse(raw)

	if err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequestWithMessage("Invalid course id"))
	}

	groupRuns, err := api.db.Course.Query().
		Where(course.ID(uuid)).
		QueryGroupRuns().
		Order(model.Desc(grouprun.FieldCreatedAt)).All(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to get group runs")
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	return c.JSON(http.StatusOK, utils.Map(groupRuns, out.ConvertModelToGroupRunOut))
}
