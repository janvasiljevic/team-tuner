package out

import (
	"jv/team-tone-tuner/model"
	"time"
)

type CourseOut struct {
	Id        string    `json:"id" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
	Code      string    `json:"code" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Colour    string    `json:"colour" validate:"required"`
}

func ConvertModelToCourseOut(m *model.Course) CourseOut {
	course := CourseOut{
		Id:        m.ID.String(),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Code:      m.Code,
		Name:      m.Name,
		Colour:    m.Colour,
	}

	return course
}

type CreatedCourseOut struct {
	Id        string    `json:"id" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
	Code      string    `json:"code" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Colour    string    `json:"colour" validate:"required"`
}

func ConvertModelToCreatedCourseOut(m *model.Course) CreatedCourseOut {
	course := CreatedCourseOut{
		Id:        m.ID.String(),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Code:      m.Code,
		Name:      m.Name,
		Colour:    m.Colour,
	}

	return course
}
