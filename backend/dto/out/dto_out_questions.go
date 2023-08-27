package out

import (
	"jv/team-tone-tuner/dto/in"
	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/model/bfiquestion"
	"jv/team-tone-tuner/utils"
)

type QuesitonOutType string

const (
	QuestionTypeExtraversion     QuesitonOutType = "extraversion"
	QuesitonOutTypeAgreeableness QuesitonOutType = "agreeableness"
	QuesitonOutTypeConscientious QuesitonOutType = "conscientious"
	QuesitonOutTypeNeuroticism   QuesitonOutType = "neuroticism"
	QuesitonOutTypeOpenness      QuesitonOutType = "openness"
)

type QuestioneItemOut struct {
	Id           string          `json:"question_id" validate:"required"`
	Question     string          `json:"question" validate:"required"`
	QuestionType QuesitonOutType `json:"question_type" validate:"required"`
	AnswerId     string          `json:"answer_id" validate:"required"`
	AnswerValue  *int            `json:"answer_value"`
}

type QuestionerOut struct {
	Title       string             `json:"title" validate:"required"`
	Description string             `json:"description" validate:"required"`
	Questions   []QuestioneItemOut `json:"questions" validate:"required"`
}

func NewQuestionerOut(modelAnswers []*model.BfiAnswer, questionType in.QuestionType) QuestionerOut {
	answersOut := utils.Map(modelAnswers, ConvertModelAnswerToOut)

	title := ""
	description := ""

	switch questionType {
	case in.QuestionTypesExtroversion:
		title = "Extroversion"
		description = "Extroversion is characterized by excitability, sociability, talkativeness, assertiveness, and high amounts of emotional expressiveness."
	case in.QuestionTypesAgreeableness:
		title = "Agreeableness"
		description = "Agreeableness is characterized by trust, altruism, kindness, and other pro-social behaviors."
	case in.QuestionTypesConscientiousness:
		title = "Conscientiousness"
		description = "Conscientiousness is characterized by thoughtfulness, good impulse control, and goal-directed behaviors."
	case in.QuestionTypesNeuroticism:
		title = "Neuroticism"
		description = "Neuroticism is characterized by sadness, moodiness, and emotional instability."
	case in.QuestionTypesOpenness:
		title = "Openness"
		description = "Openness is characterized by imagination and insight, and those high in this trait also tend to have a broad range of interests."
	}

	out := QuestionerOut{
		Title:       title,
		Description: description,
		Questions:   answersOut,
	}

	return out
}

type SubmitQuestionsOut struct {
	Success             bool               `json:"success" validate:"required"`
	UnresolvedQuestions []QuestioneItemOut `json:"unresolved_questions" validate:"required"`
}

func NewSubmitQuestionsOut(unresolvedQuestions []*model.BfiQuestion) SubmitQuestionsOut {
	if len(unresolvedQuestions) == 0 {
		return SubmitQuestionsOut{
			Success:             true,
			UnresolvedQuestions: nil,
		}
	}

	questionsOut := utils.Map(unresolvedQuestions, func(modelQ *model.BfiQuestion) QuestioneItemOut {
		answer := modelQ.Edges.BfiAnswers[0]
		return QuestioneItemOut{
			Id:          modelQ.ID.String(),
			Question:    modelQ.Questiono,
			AnswerId:    answer.ID.String(),
			AnswerValue: answer.Value,
		}
	})

	return SubmitQuestionsOut{
		Success:             false,
		UnresolvedQuestions: questionsOut,
	}
}

func ConvertModelAnswerToOut(modelAnswer *model.BfiAnswer) QuestioneItemOut {
	q := modelAnswer.Edges.BfiQuestion

	out := QuestioneItemOut{
		AnswerId:    modelAnswer.ID.String(),
		AnswerValue: modelAnswer.Value,
	}

	if q != nil {
		out.Id = q.ID.String()
		out.Question = q.Questiono

		switch q.Dimension {
		case bfiquestion.DimensionExtraversion:
			out.QuestionType = QuestionTypeExtraversion
		case bfiquestion.DimensionAgreeableness:
			out.QuestionType = QuesitonOutTypeAgreeableness
		case bfiquestion.DimensionConscientiousness:
			out.QuestionType = QuesitonOutTypeConscientious
		case bfiquestion.DimensionNeuroticism:
			out.QuestionType = QuesitonOutTypeNeuroticism
		case bfiquestion.DimensionOpenness:
			out.QuestionType = QuesitonOutTypeOpenness
		}
	}

	return out
}
