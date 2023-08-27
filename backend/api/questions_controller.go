package api

import (
	"jv/team-tone-tuner/dto/in"
	"jv/team-tone-tuner/dto/out"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

//	@Summary		Get questions
//	@Description	Get questions
//	@Tags			questions
//	@Param			request	query		in.GetQuestionsIn	true	"Get questions query"
//	@Success		200		{object}	out.QuestionerOut
//	@Router			/question [get]
func (api Api) GetQuestions(c echo.Context) error {
	var dto in.GetQuestionsIn

	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequest())
	}

	if err := c.Validate(&dto); err != nil {
		log.Warn().Err(err).Msg("Validation error")
		return c.JSON(http.StatusBadRequest, out.NewValidatorError(err))
	}

	studentId := c.Get("userId").(uuid.UUID)
	ctx := c.Request().Context()

	questions, err := api.questionStore.GetAnswersForUser(studentId, &dto.TypeOfQuestion, ctx)

	if err != nil {
		log.Error().Err(err).Msg("Error getting questions")
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	return c.JSON(http.StatusOK, out.NewQuestionerOut(questions, dto.TypeOfQuestion))
}

//	@Summary		Post questions answer
//	@Description	Post questions answer
//	@Tags			questions
//	@Param			request	body		in.PostQuestionAnswerIn	true	"Post questions answer body"
//	@Success		200		{object}	out.QuestionerOut
//	@Router			/question/answer [post]
func (api Api) PostQuestionsAnswer(c echo.Context) error {
	var dto in.PostQuestionAnswerIn

	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequest())
	}

	if err := c.Validate(&dto); err != nil {
		log.Warn().Err(err).Msg("Validation error")
		return c.JSON(http.StatusBadRequest, out.NewValidatorError(err))
	}

	userId := c.Get("userId").(uuid.UUID)

	err := api.questionStore.UpdateAnAnswer(userId, dto.AnswerID, dto.Value, c.Request().Context())

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error - Create answer")
	}

	return c.String(http.StatusOK, "OK")
}

//	@Summary		Submit questions
//	@Description	Submit questions. If all questions are answered, the BFI report will be geneareted. Else the user gets back the list of unanswered questions.
//	@Tags			questions
//	@Success		200	{object}	out.SubmitQuestionsOut
//	@Router			/question/submit [post]
func (api Api) PostSubmitQuestions(c echo.Context) error {
	userID := c.Get("userId").(uuid.UUID)

	unansweredQ, err := api.questionStore.GetUnansweredQuestions(userID, c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	if len(unansweredQ) == 0 {
		err = api.bfiReportStore.CreateBigFiveReportForUser(userID, c.Request().Context())

		if err != nil {
			log.Error().Err(err).Str("userId", userID.String()).Msg("Error creating BFI report")
			return c.JSON(http.StatusInternalServerError, out.NewInternal())
		}
	}

	return c.JSON(http.StatusOK, out.NewSubmitQuestionsOut(unansweredQ))
}
