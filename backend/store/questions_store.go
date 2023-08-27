package store

import (
	"context"
	"jv/team-tone-tuner/dto/in"
	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/model/bfianswer"
	"jv/team-tone-tuner/model/bfiquestion"
	"jv/team-tone-tuner/model/user"

	"github.com/google/uuid"
)

func (qs QuestionStore) GetAnswersForUser(userId uuid.UUID, typeOfQuestion *in.QuestionType, ctx context.Context) ([]*model.BfiAnswer, error) {
	query := qs.db.User.Query().Where(user.IDEQ(userId)).QueryBfiAnswers()

	if typeOfQuestion != nil {
		downcasted := *typeOfQuestion

		switch downcasted {
		case in.QuestionTypesExtroversion:
			query = query.Where(bfianswer.HasBfiQuestionWith(bfiquestion.DimensionEQ(bfiquestion.DimensionExtraversion)))
		case in.QuestionTypesAgreeableness:
			query = query.Where(bfianswer.HasBfiQuestionWith(bfiquestion.DimensionEQ(bfiquestion.DimensionAgreeableness)))
		case in.QuestionTypesConscientiousness:
			query = query.Where(bfianswer.HasBfiQuestionWith(bfiquestion.DimensionEQ(bfiquestion.DimensionConscientiousness)))
		case in.QuestionTypesNeuroticism:
			query = query.Where(bfianswer.HasBfiQuestionWith(bfiquestion.DimensionEQ(bfiquestion.DimensionNeuroticism)))
		case in.QuestionTypesOpenness:
			query = query.Where(bfianswer.HasBfiQuestionWith(bfiquestion.DimensionEQ(bfiquestion.DimensionOpenness)))
		}
	}

	answers, err := query.WithBfiQuestion().All(ctx)

	if err != nil {
		return nil, err
	}

	return answers, nil
}

func (qs QuestionStore) UpdateAnAnswer(userId uuid.UUID, answerId uuid.UUID, value int, ctx context.Context) error {
	_, err := qs.db.BfiAnswer.Update().Where(
		bfianswer.IDEQ(answerId),
	).SetValue(value).Save(ctx)

	return err
}

func (qs QuestionStore) GetUnansweredQuestions(userId uuid.UUID, ctx context.Context) ([]*model.BfiQuestion, error) {
	return qs.db.User.Query().
		Where(user.IDEQ(userId)).
		QueryBfiAnswers().
		Where(bfianswer.ValueIsNil()).
		QueryBfiQuestion().
		WithBfiAnswers().
		All(ctx)
}
