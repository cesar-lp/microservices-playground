package com.microservices.book.exception

import java.lang.RuntimeException

class InvalidResourceException(val invalidFields: List<FieldError>) : RuntimeException()

class ResourceNotFoundException(message: String): RuntimeException(message)