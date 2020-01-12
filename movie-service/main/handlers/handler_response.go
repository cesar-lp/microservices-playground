package handlers

import "github.com/cesar-lp/microservices-playground/movie-service/main/common"

// HandlerResponse structure.
type HandlerResponse struct {
	Service     string
	Function    string
	StatusCode  int
	Payload     interface{}
	Err         error
	FieldErrors []common.FieldError
}

// Ok builds and returns a handler response for a successful operation.
func Ok(payload interface{}) HandlerResponse {
	return HandlerResponse{
		StatusCode: 200,
		Payload:    payload,
	}
}

// Created builds and returns a handler response for a successful create operation.
func Created(payload interface{}) HandlerResponse {
	return HandlerResponse{
		StatusCode: 201,
		Payload:    payload,
	}
}

// NoContent builds and returns a handler response for a successful operation with no content returned.
func NoContent() HandlerResponse {
	return HandlerResponse{
		StatusCode: 204,
	}
}

// NotFound builds and returns a handler response for a resource not found error.
func NotFound(serviceName, functionName string, err error) HandlerResponse {
	return HandlerResponse{
		Service:    serviceName,
		Function:   functionName,
		StatusCode: 404,
		Err:        err,
	}
}

// UnprocessableEntity builds and returns a handler response for a model with invalid state.
func UnprocessableEntity(serviceName, functionName string, fieldErrors []common.FieldError) HandlerResponse {
	return HandlerResponse{
		Service:     serviceName,
		Function:    functionName,
		StatusCode:  422,
		FieldErrors: fieldErrors,
	}
}

// InternalServerError builds and returns a handler response for an internal error.
func InternalServerError(serviceName, functionName string, err error) HandlerResponse {
	return HandlerResponse{
		Service:    serviceName,
		Function:   functionName,
		StatusCode: 500,
		Err:        err,
	}
}
