package in

type CreateCourseBody struct {
	Name   string `json:"name" validate:"required,min=1,max=255"`
	Code   string `json:"code" validate:"required,min=1,max=10"`
	Colour string `json:"colour" validate:"required,hexcolor"`
}
