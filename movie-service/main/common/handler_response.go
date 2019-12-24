package common

// HandlerResponse structure
type HandlerResponse struct {
	StatusCode int
	Payload    interface{}
	Err        error
}

// Ok builds and returns a response for a successful operation
func Ok(payload interface{}) HandlerResponse {
	return buildResponse(200, payload, nil)
}

// Created builds and returns a response for an operation which successfully created a resource
func Created(payload interface{}) HandlerResponse {
	return buildResponse(201, payload, nil)
}

// NoContent builds and returns a response for a successful operation with no content returned
func NoContent() HandlerResponse {
	return buildResponse(204, nil, nil)
}

// NotFound builds and returns a response for an operation which couldn't find a resource
func NotFound(err error) HandlerResponse {
	return buildResponse(404, nil, err)
}

// InternalServerError builds and returns a response for an internal error
func InternalServerError(err error) HandlerResponse {
	return buildResponse(500, nil, err)
}

func buildResponse(statusCode int, payload interface{}, err error) HandlerResponse {
	return HandlerResponse{
		StatusCode: statusCode,
		Payload:    payload,
		Err:        err,
	}
}
