package router

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *Validator {

	validator := validator.New()

	// https://github.com/go-playground/validator/issues/900
	validator.RegisterCustomTypeFunc(ValidateUUID, uuid.UUID{})

	return &Validator{
		validator: validator,
	}
}

// ValidateUUID implements validator.CustomTypeFunc
func ValidateUUID(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(uuid.UUID); ok {
		return valuer.String()
	}
	return nil
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
