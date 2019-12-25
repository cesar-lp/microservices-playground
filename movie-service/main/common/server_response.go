package common

import (
	"encoding/json"
	"net/http"
	"time"
)

type baseResponse struct {
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

type validationResponse struct {
	Message     string       `json:"message"`
	FieldErrors []FieldError `json:"fieldErrors"`
	Path        string       `json:"path"`
	Timestamp   time.Time    `json:"timestamp"`
}

// FieldError structure
type FieldError struct {
	FieldName    string `json:"fieldName"`
	Error        string `json:"error"`
	InvalidValue string `json:"invalidValue"`
}

// ServerResponse builds a HTTP response based on a handler response
func ServerResponse(w http.ResponseWriter, r *http.Request, hr HandlerResponse) {
	switch hr.StatusCode {
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(internalServerError(r.RequestURI, hr.Err.Error()))
	case 422:
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(validationError(r.RequestURI, hr.FieldErrors))
	case 404:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resourceNotFound(r.RequestURI, hr.Err.Error()))
	case 204:
		w.WriteHeader(http.StatusNoContent)
	case 201:
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(hr.Payload)
	default:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hr.Payload)
	}
}

func internalServerError(path string, message string) baseResponse {
	return baseResponse{
		Error:     "Internal Server Error",
		Message:   message,
		Path:      path,
		Timestamp: time.Now(),
	}
}

func validationError(path string, fieldErrors []FieldError) validationResponse {
	return validationResponse{
		Message:     "Validation Failed",
		FieldErrors: fieldErrors,
		Path:        path,
		Timestamp:   time.Now(),
	}
}

func resourceNotFound(path string, message string) baseResponse {
	return baseResponse{
		Error:     "Resource Not Found",
		Message:   message,
		Path:      path,
		Timestamp: time.Now(),
	}
}
