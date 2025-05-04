package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse represents a standard JSON response structure for errors and messages
type JSONResponse struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

// RespondWithJSON sends a JSON response with the given status code
func RespondWithJSON(writer http.ResponseWriter, statusCode int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(data)
}

// RespondWithError sends a JSON error response
func RespondWithError(writer http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(writer, statusCode, JSONResponse{
		Error: message,
	})
}

// RespondWithSuccess sends a JSON success response with data
func RespondWithSuccess(writer http.ResponseWriter, statusCode int, data any) {
	RespondWithJSON(writer, statusCode, data)
}

// RespondWithMessage sends a JSON response with a message
func RespondWithMessage(writer http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(writer, statusCode, JSONResponse{
		Message: message,
	})
}
