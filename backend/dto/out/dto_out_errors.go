package out

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}

func NewValidatorError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		e.Errors[v.Field()] = fmt.Sprintf("%v", v.Tag())
	}
	return e
}

func BadRequest() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "bad request"
	return e
}

func BadRequestUUID() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "bad request: invalid uuid"
	return e
}

func BadRequestWithMessage(message string) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = message
	return e
}

func AccessForbidden() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "access forbidden"
	return e
}

func NotFound() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "resource not found"
	return e
}

func NewInternal() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "internal server error"
	return e
}
