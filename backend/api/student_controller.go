package api

import (
	"jv/team-tone-tuner/dto/in"
	"jv/team-tone-tuner/dto/out"
	"jv/team-tone-tuner/model/bfireport"
	"jv/team-tone-tuner/model/course"
	"jv/team-tone-tuner/model/user"
	"jv/team-tone-tuner/utils"
	"net/http"

	"entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

//	@Summary		Get students
//	@Description	Get students
//	@Tags			students
//	@Param			request	query		in.GetStudentsQuery							true	"Get students query"
//	@Success		200		{object}	out.PaginationOut{content=[]out.StudentOut}	"Students filtered by query"
//	@Router			/student [get]
func (api Api) GetStudents(c echo.Context) error {
	ctx := c.Request().Context()
	var dto in.GetStudentsQuery

	if err := c.Bind(&dto); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	query := api.db.User.Query()

	page := 0
	pageSize := 10

	if dto.Page != nil {
		page = *dto.Page
	}

	if dto.PageSize != nil {
		pageSize = *dto.PageSize
	}

	if dto.CourseId != nil {
		query = query.Where(
			user.HasCoursesWith(course.IDEQ(*dto.CourseId)),
		)
	}

	if dto.CompletedQuestioner != nil {
		if *dto.CompletedQuestioner {
			query = query.Where(user.HasBfiReport())
		} else {
			query = query.Where(user.Not(user.HasBfiAnswers()))
		}
	}

	if dto.SortField != nil {
		order := sql.OrderDesc()

		if dto.SortOrder != nil {
			if *dto.SortOrder == "ASC" {
				order = sql.OrderAsc()
			}
		}

		switch *dto.SortField {
		case (in.StudentSortFieldOpenness):
			query = query.Order(user.ByBfiReportField(bfireport.FieldOpenness, order))
		case (in.StudentSortFieldConscientious):
			query = query.Order(user.ByBfiReportField(bfireport.FieldConscientiousness, order))
		case (in.StudentSortFieldExtraversion):
			query = query.Order(user.ByBfiReportField(bfireport.FieldExtraversion, order))
		case (in.StudentSortFieldAgreeableness):
			query = query.Order(user.ByBfiReportField(bfireport.FieldAgreeableness, order))
		case (in.StudentSortFieldNeuroticism):
			query = query.Order(user.ByBfiReportField(bfireport.FieldNeuroticism, order))
		default:
			return c.String(http.StatusBadRequest, "Bad request - Unknown sort field")
		}
	} else {
		query = query.Order(user.ByGithubUsername(sql.OrderAsc()))
	}

	clonedQuery := query.Clone()

	students, err := clonedQuery.Limit(pageSize).WithBfiReport().Offset(page * pageSize).All(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to get students")
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	totalStudentCount, err := query.Count(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to get students count")
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	studentsOut := utils.Map(students, out.ConvertModelToStudentOut)

	pagination := out.NewPaginationOut(totalStudentCount, page, pageSize, studentsOut)

	return c.JSON(http.StatusOK, pagination)
}
