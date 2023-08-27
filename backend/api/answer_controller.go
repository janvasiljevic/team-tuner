package api

import (
	"jv/team-tone-tuner/dto/out"
	"jv/team-tone-tuner/utils"

	"jv/team-tone-tuner/model/bfianswer"
	"jv/team-tone-tuner/model/bfireport"
	"jv/team-tone-tuner/model/user"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

//	@Summary		Get all answers
//	@Description	Get all answers: Get all answers for the current user, based on the JWT
//	@Tags			answer
//	@Success		200	{array}	out.QuestioneItemOut
//	@Router			/answer [get]
func (api Api) GetAllAnswers(c echo.Context) error {

	userId := c.Get("userId").(uuid.UUID)

	answers, err := api.questionStore.GetAnswersForUser(userId, nil, c.Request().Context())

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	out := utils.Map(answers, out.ConvertModelAnswerToOut)

	return c.JSON(http.StatusOK, out)
}

//	@Summary		Get questioner status
//	@Description	Get questioner status: Get the status of the questioner for the current user, based on the JWT
//	@Tags			answer
//	@Success		200	{object}	out.QuestionerStatus
//	@Router			/answer/questioner-status [get]
func (api Api) GetQuestionerStatus(c echo.Context) error {

	userId := c.Get("userId").(uuid.UUID)

	// Get the count of unasweredCount that are nil
	unasweredCount, err := api.db.User.Query().Where(user.ID(userId)).QueryBfiAnswers().Where(bfianswer.ValueIsNil()).Count(c.Request().Context())

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, out.NewQuestionerStatus(unasweredCount == 0))
}

//	@Summary		Get BFI report
//	@Description	Get BFI report
//	@Tags			answer
//	@Success		200	{object}	out.BfiReportOut
//	@Router			/answer/bfi-report [get]
func (api Api) GetBfiReport(c echo.Context) error {

	userId := c.Get("userId").(uuid.UUID)

	report, err := api.db.BfiReport.Query().Where(bfireport.HasStudentWith(user.IDEQ(userId))).Only(c.Request().Context())

	if err != nil {
		log.Error().Err(err).Str("userId", userId.String()).Msg("Error while getting BFI report")
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	if report == nil {
		log.Warn().Str("userId", userId.String()).Msg("Didn't find BFI report")
		return c.JSON(http.StatusNotFound, out.NotFound())
	}

	return c.JSON(http.StatusOK, out.ConvertModelToBfiReportOut(report))
}
