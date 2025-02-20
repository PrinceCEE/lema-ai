package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var v = validator.New(validator.WithRequiredStructEnabled())

func ValidateData(data any) map[string]any {
	err := v.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return map[string]any{"error": err.Error()}
		}

		errors := make(map[string]any)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Error()
		}

		return errors
	}

	return nil
}

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
