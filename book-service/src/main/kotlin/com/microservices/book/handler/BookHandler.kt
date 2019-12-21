package com.microservices.book.handler

import com.microservices.book.converter.toBook
import com.microservices.book.entity.Book
import com.microservices.book.exception.ResourceNotFoundException
import com.microservices.book.repository.BookRepository
import com.microservices.book.request.PersistBookRequest
import com.microservices.book.validation.ValidationHandler
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.server.ServerRequest
import org.springframework.web.reactive.function.server.ServerResponse
import org.springframework.web.reactive.function.server.ServerResponse.created
import org.springframework.web.reactive.function.server.ServerResponse.noContent
import org.springframework.web.reactive.function.server.ServerResponse.notFound
import org.springframework.web.reactive.function.server.ServerResponse.ok
import org.springframework.web.reactive.function.server.body
import org.springframework.web.reactive.function.server.bodyToMono
import reactor.core.publisher.Mono
import reactor.core.publisher.switchIfEmpty
import java.net.URI

@Component
class BookHandler(
	private val validator: ValidationHandler,
	private val repository: BookRepository
) {

	// TODO: cleanup?

	private val baseUri = "/books"

	fun getAllBooks(req: ServerRequest) =
		ok().body(repository.findAll())

	fun saveBook(req: ServerRequest) =
		validator.validate(req.bodyToMono<PersistBookRequest>())
			.flatMap { repository.save(it.toBook()) }
			.flatMap { created(URI.create("$baseUri/${it.id}")).body(Mono.just(it)) }

	fun findBook(req: ServerRequest): Mono<ServerResponse> {
		val id = req.pathVariable("id")
		return repository
			.findById(id)
			.switchIfEmpty(Mono.error(ResourceNotFoundException("Book not found for id $id")))
			.flatMap { ok().body<Book>(Mono.just(it)) }
	}

	fun updateBook(req: ServerRequest) =
		validator.validate(req.bodyToMono<PersistBookRequest>())
			.map { it.toBook() }
			.flatMap { updated ->
				repository.findById(req.pathVariable("id"))
					.flatMap { existing ->
						existing.ranking = updated.ranking
						repository.save(existing)
					}
			}.flatMap { ok().body(Mono.just(it)) }

	fun deleteBook(req: ServerRequest): Mono<ServerResponse> {
		val id = req.pathVariable("id")
		return repository.findById(id)
			.switchIfEmpty(Mono.error(ResourceNotFoundException("Book not found for id $id")))
			.flatMap { repository.delete(it).then(Mono.just(it)) }
			.flatMap { noContent().build() }
	}
}