package out

import (
	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/utils"
	"time"
)

type GroupOut struct {
	Id        string    `json:"id" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updateAt" validate:"required"`

	CourseCode   string `json:"courseCode" validate:"required"`
	CourseName   string `json:"courseName" validate:"required"`
	CourseColour string `json:"courseColour" validate:"required"`

	Students []StudentOut `json:"students" validate:"required"`
	Name     string       `json:"name" validate:"required"`
}

type DetailedGroupOut struct {
	GroupOut

	BigFiveBoxPlot BigFiveBoxPlot `json:"bigFiveBoxPlot" validate:"required"`
}

func ConvertModelToGroupOut(m *model.Group) GroupOut {
	group := GroupOut{
		Id:        m.ID.String(),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Students:  utils.Map(m.Edges.Students, ConvertModelToStudentOut),
		Name:      m.Name,
	}

	course := m.Edges.Course

	if course != nil {
		group.CourseCode = course.Code
		group.CourseName = course.Name
		group.CourseColour = course.Colour
	}

	return group
}

func NewDetailedGroupOut(m *model.Group) DetailedGroupOut {
	reports := utils.Map(m.Edges.Students, func(s *model.User) *model.BfiReport {
		return s.Edges.BfiReport
	})

	return DetailedGroupOut{
		GroupOut:       ConvertModelToGroupOut(m),
		BigFiveBoxPlot: NewBigFiveBoxPlotFromReports(reports),
	}
}
