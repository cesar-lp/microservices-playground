package handlers

import "github.com/cesar-lp/microservices-playground/movie-service/main/common"

// HandlerResponse structure
type HandlerResponse struct {
	StatusCode  int
	Payload     interface{}
	Err         error
	FieldErrors []common.FieldError
}

// Ok builds and returns a response for a successful operation
func Ok(payload interface{}) HandlerResponse {
	return HandlerResponse{
		StatusCode: 200,
		Payload:    payload,
	}
}

// Created builds and returns a response for an operation which successfully created a resource
func Created(payload interface{}) HandlerResponse {
	return HandlerResponse{
		StatusCode: 201,
		Payload:    payload,
	}
}

// NoContent builds and returns a response for a successful operation with no content returned
func NoContent() HandlerResponse {
	return HandlerResponse{
		StatusCode: 204,
	}
}

// NotFound builds and returns a response for an operation which couldn't find a resource
func NotFound(err error) HandlerResponse {
	return HandlerResponse{
		StatusCode: 404,
		Err:        err,
	}
}

func UnprocessableEntity(fieldErrors []common.FieldError) HandlerResponse {
	return HandlerResponse{
		StatusCode:  422,
		FieldErrors: fieldErrors,
	}
}

// InternalServerError builds and returns a response for an internal error
func InternalServerError(err error) HandlerResponse {
	return HandlerResponse{
		StatusCode: 500,
		Err:        err,
	}
}
