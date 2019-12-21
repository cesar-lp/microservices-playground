package com.microservices.book.routing

import com.microservices.book.handler.BookHandler
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.MediaType
import org.springframework.web.reactive.function.server.router

@Configuration
class BookRouter(private val handler: BookHandler) {

	@Bean
	fun routes() = router {
		"/api/books".nest {
			accept(MediaType.APPLICATION_JSON).and(contentType(MediaType.APPLICATION_JSON)).nest {
				GET("", handler::getAllBooks)
				POST("", handler::saveBook)

				"/{id}".nest {
					GET("", handler::findBook)
					PUT("", handler::updateBook)
					DELETE("", handler::deleteBook)
				}
			}
		}
	}
}