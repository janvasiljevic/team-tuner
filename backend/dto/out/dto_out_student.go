package out

import (
	"jv/team-tone-tuner/model"
	"time"
)

type StudentOut struct {
	Id                    string     `json:"id" validate:"required"`
	CreatedAt             time.Time  `json:"createdAt" validate:"required"`
	UpdatedAt             time.Time  `json:"updatedAt" validate:"required"`
	GithubUsername        string     `json:"githubUsername" validate:"required"`
	UniveresityId         *string    `json:"univeresityId,omitempty"`
	CompletedQuestioner   bool       `json:"completedQuestioner" validate:"required"`
	CompletedQuestionerAt *time.Time `json:"completedQuestionerAt,omitempty"`
	Concientiousness      *float64   `json:"concientiousness,omitempty"`
	Openness              *float64   `json:"openness,omitempty"`
	Extraversion          *float64   `json:"extraversion,omitempty"`
	Agreeableness         *float64   `json:"agreeableness,omitempty"`
	Neuroticism           *float64   `json:"neuroticism,omitempty"`
}

func ConvertModelToStudentOut(m *model.User) StudentOut {
	student := StudentOut{
		Id:             m.ID.String(),
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		GithubUsername: m.GithubUsername,
		UniveresityId:  m.UniveresityID,
	}

	report := m.Edges.BfiReport

	if report != nil {
		student.CompletedQuestioner = true
		student.CompletedQuestionerAt = &report.CreatedAt
		student.Concientiousness = &report.Conscientiousness.PointsNormalized
		student.Openness = &report.Openness.PointsNormalized
		student.Extraversion = &report.Extraversion.PointsNormalized
		student.Agreeableness = &report.Agreeableness.PointsNormalized
		student.Neuroticism = &report.Neuroticism.PointsNormalized
	} else {
		student.CompletedQuestioner = false
	}

	return student
}
