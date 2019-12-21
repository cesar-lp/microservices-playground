package com.microservices.book.validation

import com.microservices.book.exception.FieldError
import com.microservices.book.exception.InvalidResourceException
import org.springframework.stereotype.Component
import reactor.core.publisher.Mono
import javax.validation.Validator
import kotlin.streams.toList

@Component
class ValidationHandler(val validator: Validator) {

	fun <T> validate(toValidate: Mono<T>): Mono<T> {
		return toValidate.flatMap { reqObj ->
			val errors = validator.validate(reqObj)

			if (errors.isNotEmpty()) {
				val fieldErrors = errors.stream()
					.map { FieldError(it.propertyPath.toString(), it.message, it.invalidValue.toString()) }
					.toList()

				throw InvalidResourceException(fieldErrors)
			}

			Mono.just(reqObj)
		}
	}
}