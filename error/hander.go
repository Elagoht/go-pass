package errorHandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

/* Represents a validation error response */
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

/* Returns a user-friendly error message for validation errors */
func getValidationErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Must be at least " + err.Param() + " characters"
	case "max":
		return "Must be at most " + err.Param() + " characters"
	case "url":
		return "Must be a valid URL"
	default:
		return "Invalid value"
	}
}

/* Gets the first validation error and returns a proper response */
func HandleValidationError(writer http.ResponseWriter, err error) {
	validationErrors := err.(validator.ValidationErrors)
	if len(validationErrors) > 0 {
		firstError := validationErrors[0]
		errorResponse := ValidationError{
			Field:   firstError.Field(),
			Message: getValidationErrorMessage(firstError),
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorResponse)
	}
}
