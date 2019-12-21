package com.microservices.book.exception

import java.time.LocalDateTime

class InternalServerErrorResponse(
	val message: String = "Internal Server Error",
	val statusCode: Int = 500,
	val path: String,
	val timestamp: LocalDateTime = LocalDateTime.now()
)

class ResourceNotFoundResponse(
	val message: String,
	val statusCode: Int = 404,
	val path: String,
	val timestamp: LocalDateTime = LocalDateTime.now()
)

class InvalidResourceResponse(
	val message: String = "Invalid Resource Body",
	val statusCode: Int = 400,
	val fieldErrors: List<FieldError> = emptyList(),
	val path: String,
	val timestamp: LocalDateTime = LocalDateTime.now()
)

class FieldError(val fieldName: String, val message: String, val invalidValue: String)
