package com.microservices.book.routing

import com.microservices.book.handler.BookHandler
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.web.reactive.function.server.router

@Configuration
class BookRoutes(private val handler: BookHandler) {

	@Bean
	fun routes() = router {
		"/api/books".nest {
			GET("", handler::getAllBooks)
			POST("", handler::saveBook)

			"/{id}".nest {
				GET("", handler::getBookById)
				PUT("", handler::updateBook)
				DELETE("", handler::deleteBook)
			}
		}
	}
}