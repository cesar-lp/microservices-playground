package com.microservices.book.exception

import com.fasterxml.jackson.databind.ObjectMapper
import org.springframework.core.annotation.Order
import org.springframework.http.HttpStatus
import org.springframework.http.MediaType
import org.springframework.stereotype.Component
import org.springframework.web.server.ServerWebExchange
import org.springframework.web.server.WebExceptionHandler
import reactor.core.publisher.Mono

@Component
@Order(-2)
class GlobalExceptionHandler(private val objectMapper: ObjectMapper) : WebExceptionHandler {

	override fun handle(exchange: ServerWebExchange, ex: Throwable): Mono<Void> {
		val path = exchange.request.path.value()

		// TODO: cleanup
		val errorObj = when (ex) {
			is InvalidResourceException -> {
				exchange.response.statusCode = HttpStatus.BAD_REQUEST
				InvalidResourceResponse(fieldErrors = ex.invalidFields, path = path)
			}
			is ResourceNotFoundException -> {
				exchange.response.statusCode = HttpStatus.NOT_FOUND
				ResourceNotFoundResponse(message = ex.localizedMessage, path = path)
			}
			else -> {
				exchange.response.statusCode = HttpStatus.INTERNAL_SERVER_ERROR
				InternalServerErrorResponse(path = path)
			}
		}

		val bytes = objectMapper.writeValueAsBytes(errorObj)
		val buffer = exchange.response.bufferFactory().wrap(bytes)
		exchange.response.headers.contentType = MediaType.APPLICATION_JSON
		return exchange.response.writeWith(Mono.just(buffer))
	}
}