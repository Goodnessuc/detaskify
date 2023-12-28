package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (cv *Validator) Validate(i interface{}) error {
	if cv.validator == nil {
		return errors.New("validator is not initialized")
	}
	return cv.validator.Struct(i)
}
