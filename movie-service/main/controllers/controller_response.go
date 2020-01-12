package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cesar-lp/microservices-playground/movie-service/main/common"
	"github.com/cesar-lp/microservices-playground/movie-service/main/handlers"

	log "github.com/sirupsen/logrus"
)

// BaseErrorResponse structure.
type BaseErrorResponse struct {
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

// ValidationErrorResponse structure.
type ValidationErrorResponse struct {
	Message     string              `json:"message"`
	FieldErrors []common.FieldError `json:"fieldErrors"`
	Path        string              `json:"path"`
	Timestamp   time.Time           `json:"timestamp"`
}

// ServerResponse builds a HTTP response based on a handler response.
func ServerResponse(w http.ResponseWriter, r *http.Request, hr handlers.HandlerResponse) {
	switch hr.StatusCode {
	case 500:
		logError(hr)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(internalServerError(r.URL.Path, hr.Err.Error()))
	case 422:
		logError(hr)
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ValidationError(r.URL.Path, hr.FieldErrors))
	case 404:
		logError(hr)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ResourceNotFound(r.URL.Path, hr.Err.Error()))
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

// ValidationError builds and returns a server response object for a validation error.
func ValidationError(path string, fieldErrors []common.FieldError) ValidationErrorResponse {
	return ValidationErrorResponse{
		Message:     "Validation Failed",
		FieldErrors: fieldErrors,
		Path:        path,
		Timestamp:   time.Now(),
	}
}

// ResourceNotFound builds and returns a server response object for a resource not found error.
func ResourceNotFound(path string, message string) BaseErrorResponse {
	return BaseErrorResponse{
		Error:     "Resource Not Found",
		Message:   message,
		Path:      path,
		Timestamp: time.Now(),
	}
}

func internalServerError(path string, message string) BaseErrorResponse {
	return BaseErrorResponse{
		Error:     "Internal Server Error",
		Message:   message,
		Path:      path,
		Timestamp: time.Now(),
	}
}

func logError(hr handlers.HandlerResponse) {
	log.Errorf("(%s) - %s: %s", hr.Service, hr.Function, hr.Err.Error())
}
