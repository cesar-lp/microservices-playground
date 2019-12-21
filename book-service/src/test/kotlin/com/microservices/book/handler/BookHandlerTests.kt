package com.microservices.book.handler

import com.fasterxml.jackson.databind.ObjectMapper
import com.microservices.book.entity.Book
import com.microservices.book.repository.BookRepository
import com.microservices.book.routing.BookRouter
import io.mockk.clearAllMocks
import io.mockk.confirmVerified
import io.mockk.every
import io.mockk.impl.annotations.MockK
import io.mockk.junit5.MockKExtension
import io.mockk.verify
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.BeforeAll
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.TestInstance
import org.junit.jupiter.api.extension.ExtendWith
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.MediaType
import org.springframework.test.web.reactive.server.WebTestClient
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono

@ExtendWith(MockKExtension::class)
@TestInstance(TestInstance.Lifecycle.PER_CLASS)
internal class BookHandlerTests {

	companion object {
		val objectMapper = ObjectMapper()
	}

	@MockK
	private lateinit var bookRepository: BookRepository

	@Autowired
	private lateinit var webClient: WebTestClient

	private val baseUri = "/api/books"

	@BeforeAll
	fun setUp() {
		val router = BookRouter(BookHandler(bookRepository))
		webClient = WebTestClient.bindToRouterFunction(router.routes()).build()
	}

	@AfterEach
	fun afterEach() {
		clearAllMocks()
	}

	@Test
	fun `get all books`() {
		val books = arrayOf(
			Book("1", "A Dance with Dragons", 1, "George R. R. Martin", "978-0345540560"),
			Book("2", "The Winds of Winter", -1, "George R. R. Martin", "978-0345540561")
		)

		every { bookRepository.findAll() } returns Flux.fromArray(books)

		webClient.get().uri(baseUri)
			.accept(MediaType.APPLICATION_JSON)
			.header("Content-type", "application/json")
			.exchange()
			.expectStatus().isOk
			.expectBody().json(objectMapper.writeValueAsString(books))

		verify(exactly = 1) { bookRepository.findAll() }
		confirmVerified(bookRepository)
	}

	@Test
	fun `create book`() {
		val bookToCreate =
			Book(name = "A Dream of Spring", ranking = -1, author = "George R. R. Martin", isbn = "978-0345540561")
		val createdBook = bookToCreate.copy(id = "10")

		every { bookRepository.save(bookToCreate) } returns Mono.just(createdBook)

		webClient.post().uri(baseUri)
			.accept(MediaType.APPLICATION_JSON)
			.header("Content-type", "application/json")
			.bodyValue(bookToCreate)
			.exchange()
			.expectStatus().isCreated
			.expectBody().json(objectMapper.writeValueAsString(createdBook))

		verify(exactly = 1) { bookRepository.save(bookToCreate) }
		confirmVerified(bookRepository)
	}

	@Test
	fun `find book`() {
		val id = "1"
		val book = Book(id, "A Dance with Dragons", 1, "George R. R. Martin", "978-0345540560")

		every { bookRepository.findById(id) } returns Mono.just(book)

		webClient.get().uri("$baseUri/$id")
			.accept(MediaType.APPLICATION_JSON)
			.header("Content-type", "application/json")
			.exchange()
			.expectStatus().isOk
			.expectBody().json(objectMapper.writeValueAsString(book))

		verify(exactly = 1) { bookRepository.findById(id) }
		confirmVerified(bookRepository)
	}

	@Test
	fun `update book`() {
		val id = "1"
		val existingBook =
			Book(name = "A Dream of Spring", ranking = -1, author = "George R. R. Martin", isbn = "978-0345540561")
		val updatedBook = existingBook.copy(ranking = 5)

		every { bookRepository.findById(id) } returns Mono.just(existingBook)
		every { bookRepository.save(updatedBook) } returns Mono.just(updatedBook)

		webClient.put().uri("$baseUri/$id")
			.accept(MediaType.APPLICATION_JSON)
			.header("Content-type", "application/json")
			.bodyValue(updatedBook)
			.exchange()
			.expectStatus().isOk
			.expectBody().json(objectMapper.writeValueAsString(updatedBook))

		verify(exactly = 1) {
			bookRepository.save(updatedBook)
			bookRepository.findById(id)
		}
		confirmVerified(bookRepository)
	}

	@Test
	fun `delete book`() {
		val id = "1"

		every { bookRepository.deleteById(id) } returns Mono.empty()

		webClient.delete().uri("$baseUri/$id")
			.accept(MediaType.APPLICATION_JSON)
			.header("Content-type", "application/json")
			.exchange()
			.expectStatus().isNoContent

		verify(exactly = 1) { bookRepository.deleteById(id) }
		confirmVerified(bookRepository)
	}
}