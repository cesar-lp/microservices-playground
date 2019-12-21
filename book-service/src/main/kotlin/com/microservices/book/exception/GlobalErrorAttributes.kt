package com.microservices.book.exception

import org.springframework.boot.web.reactive.error.DefaultErrorAttributes
import org.springframework.http.HttpStatus
import org.springframework.web.reactive.function.server.ServerRequest

class GlobalErrorAttributes : DefaultErrorAttributes() {

	override fun getErrorAttributes(request: ServerRequest?, includeStackTrace: Boolean): MutableMap<String, Any> {
		val map = super.getErrorAttributes(request, includeStackTrace)
		map["status"] = HttpStatus.BAD_REQUEST
		map["message"] = "Invalid entity"
		return map
	}
}