package out

import "jv/team-tone-tuner/model"

type BfiReportOut struct {
	Openness          float64 `json:"openness" validate:"required"`
	Conscientiousness float64 `json:"conscientiousness" validate:"required"`
	Extraversion      float64 `json:"extraversion" validate:"required"`
	Agreeableness     float64 `json:"agreeableness" validate:"required"`
	Neuroticism       float64 `json:"neuroticism" validate:"required"`
}

func ConvertModelToBfiReportOut(m *model.BfiReport) BfiReportOut {
	return BfiReportOut{
		Openness:          m.Openness.PointsNormalized,
		Conscientiousness: m.Conscientiousness.PointsNormalized,
		Extraversion:      m.Extraversion.PointsNormalized,
		Agreeableness:     m.Agreeableness.PointsNormalized,
		Neuroticism:       m.Neuroticism.PointsNormalized,
	}
}
