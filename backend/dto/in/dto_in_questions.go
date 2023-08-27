package in

import "github.com/google/uuid"

type QuestionType string

const (
	QuestionTypesExtroversion      QuestionType = "extroversion"
	QuestionTypesAgreeableness     QuestionType = "agreeableness"
	QuestionTypesConscientiousness QuestionType = "conscientiousness"
	QuestionTypesNeuroticism       QuestionType = "neuroticism"
	QuestionTypesOpenness          QuestionType = "openness"
)

type GetQuestionsIn struct {
	TypeOfQuestion QuestionType `query:"typeOfQuestion" validate:"required"`
}

type PostQuestionAnswerIn struct {
	Value    int       `body:"value" validate:"required"`
	AnswerID uuid.UUID `body:"answerId" validate:"required,uuid4"`
}
