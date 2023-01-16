package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/kryz81/go_api_and_react/types"
)

func ExtractValidationErrors(validationErrors validator.ValidationErrors) []types.BodyError {
	errors := make([]types.BodyError, len(validationErrors))
	for i, e := range validationErrors {
		errors[i] = types.BodyError{
			Field:   e.Field(),
			Message: e.Error(),
		}
	}
	return errors
}
