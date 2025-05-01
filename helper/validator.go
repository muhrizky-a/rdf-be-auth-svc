package helper

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/ryakadev/rdf-be-auth-svc/exceptions"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	validatorCustom := &Validator{}

	v := validator.New()

	// register function to get tag name from json tags.
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	validatorCustom.validator = v

	return validatorCustom
}

func (v *Validator) CreateValidationError(err error) (*exceptions.ValidationError, error) {
	validationErrors := &exceptions.ValidationError{}

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			scope := &exceptions.FailedValidation{
				Field:      e.Field(),
				Tag:        e.Tag(),
				ParamValue: e.Param(),
			}
			validationErrors.Validations = append(validationErrors.Validations, scope)
		}
	} else {
		return nil, errors.New("VALIDATOR.UNKNOWN")
	}

	return validationErrors, nil
}

func (v *Validator) Validate(i any) error {
	return v.validator.Struct(i)
}
