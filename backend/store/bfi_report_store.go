package store

import (
	"context"
	"errors"
	"jv/team-tone-tuner/model/bfianswer"
	"jv/team-tone-tuner/model/bfiquestion"
	"jv/team-tone-tuner/model/bfireport"
	"jv/team-tone-tuner/model/user"
	"jv/team-tone-tuner/schema"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (bs BfiReportStore) CreateBigFiveReportForUser(userId uuid.UUID, ctx context.Context) error {

	// check if the user already has a reportExists
	reportExists, err := bs.db.BfiReport.Query().Where(bfireport.HasStudentWith(user.IDEQ(userId))).Exist(ctx)

	if err != nil {
		return err
	}

	if reportExists {
		log.Warn().Str("userId", userId.String()).Msgf("User %s already has a report", userId)
		return errors.New("User already has a report")
	}

	answers, err := bs.db.BfiQuestion.Query().QueryBfiAnswers().Where(
		bfianswer.HasStudentWith(user.IDEQ(userId))).WithBfiQuestion().All(ctx)

	if err != nil {
		return err
	}

	// big five
	var extraversion, agreeableness, conscientiousness, neuroticism, openness schema.BfiReportItem

	for i, answer := range answers {
		// This should never happen
		if answer.Value == nil {
			log.Error().Str("userId", userId.String()).Msgf("Answer %d has no value", i)
			continue
		}

		q := answer.Edges.BfiQuestion

		if q.Dimension == bfiquestion.DimensionExtraversion {
			incrmBigFiveItem(&extraversion, *answer.Value)
		}

		if q.Dimension == bfiquestion.DimensionAgreeableness {
			incrmBigFiveItem(&agreeableness, *answer.Value)
		}

		if q.Dimension == bfiquestion.DimensionConscientiousness {
			incrmBigFiveItem(&conscientiousness, *answer.Value)
		}

		if q.Dimension == bfiquestion.DimensionNeuroticism {
			incrmBigFiveItem(&neuroticism, *answer.Value)
		}

		if q.Dimension == bfiquestion.DimensionOpenness {
			incrmBigFiveItem(&openness, *answer.Value)
		}
	}

	extraversion.PointsNormalized = generateNormalizedPoints(extraversion)
	agreeableness.PointsNormalized = generateNormalizedPoints(agreeableness)
	conscientiousness.PointsNormalized = generateNormalizedPoints(conscientiousness)
	neuroticism.PointsNormalized = generateNormalizedPoints(neuroticism)
	openness.PointsNormalized = generateNormalizedPoints(openness)

	_, err = bs.db.BfiReport.Create().
		SetStudentID(userId).
		SetExtraversion(extraversion).
		SetAgreeableness(agreeableness).
		SetConscientiousness(conscientiousness).
		SetNeuroticism(neuroticism).
		SetOpenness(openness).
		Save(ctx)

	return err
}

func incrmBigFiveItem(item *schema.BfiReportItem, value int) {
	item.PointsSum += value
	item.PointsMax += 5
	item.PointsMin += 1
}

func generateNormalizedPoints(item schema.BfiReportItem) float64 {
	if item.PointsMax-item.PointsMin == 0 {
		return 0
	}

	return float64(item.PointsSum-item.PointsMin) / float64(item.PointsMax-item.PointsMin)
}
