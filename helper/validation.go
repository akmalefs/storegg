package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errors []string

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return errors
	}

	for _, e := range validationErrors {
		message := fmt.Sprintf("Field %s on field %s", e.Field(), e.Tag())
		errors = append(errors, message)
	}
	return errors
}
