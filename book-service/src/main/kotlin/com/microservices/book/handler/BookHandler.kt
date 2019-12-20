package com.microservices.book.handler

import com.microservices.book.entity.Book
import com.microservices.book.repository.BookRepository
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.server.ServerRequest
import org.springframework.web.reactive.function.server.ServerResponse.created
import org.springframework.web.reactive.function.server.ServerResponse.noContent
import org.springframework.web.reactive.function.server.ServerResponse.ok
import org.springframework.web.reactive.function.server.body
import org.springframework.web.reactive.function.server.bodyToMono
import reactor.core.publisher.Mono
import java.net.URI

@Component
class BookHandler(private val repository: BookRepository) {

	fun getAllBooks(req: ServerRequest) =
		ok().body(repository.findAll())

	fun getBookById(req: ServerRequest) =
		repository
			.findById(req.pathVariable("id"))
			.flatMap { ok().body<Book>(Mono.just(it)) }

	fun saveBook(req: ServerRequest) =
		req.bodyToMono<Book>()
			.flatMap { repository.save(it) }
			.flatMap { created(URI.create("/items/${it.id}")).body(Mono.just(it)) }

	fun updateBook(req: ServerRequest) =
		repository.findById(req.pathVariable("id"))
			.flatMap { existing ->
				req.bodyToMono<Book>().flatMap { updated ->
					existing.ranking = updated.ranking
					ok().body(repository.save(existing))
				}
			}

	fun deleteBook(req: ServerRequest) =
		repository.deleteById(req.pathVariable("id"))
			.flatMap { noContent().build() }
}