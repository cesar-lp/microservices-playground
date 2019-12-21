package com.microservices.book.handler

import com.fasterxml.jackson.databind.ObjectMapper
import com.microservices.book.converter.toBook
import com.microservices.book.entity.Book
import com.microservices.book.repository.BookRepository
import com.microservices.book.request.PersistBookRequest
import com.microservices.book.routing.BookRouter
import com.microservices.book.validation.ValidationHandler
import io.mockk.clearAllMocks
import io.mockk.confirmVerified
import io.mockk.every
import io.mockk.junit5.MockKExtension
import io.mockk.mockk
import io.mockk.verify
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.BeforeAll
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.TestInstance
import org.junit.jupiter.api.extension.ExtendWith
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.autoconfigure.web.reactive.WebFluxTest
import org.springframework.boot.test.mock.mockito.MockBean
import org.springframework.http.MediaType
import org.springframework.test.web.reactive.server.WebTestClient
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono

@ExtendWith(MockKExtension::class)
@TestInstance(TestInstance.Lifecycle.PER_CLASS)
internal class BookHandlerTests {

	// TODO: figure out how to run tests correctly

	companion object {
		val objectMapper = ObjectMapper()
	}

	@MockBean
	private lateinit var bookRepository: BookRepository // = mockk()

	private val validationHandler: ValidationHandler = mockk()

	@Autowired
	private lateinit var webClient: WebTestClient

	@BeforeAll
	fun setUp() {
		val router = BookRouter(BookHandler(validationHandler, bookRepository))
		webClient = WebTestClient
			.bindToRouterFunction(router.routes())
			.configureClient()
			.baseUrl("/api/books")
			.build()
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

		webClient.get().uri("")
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
		val persistBookRequest = PersistBookRequest(
			name = "A Dream of Spring", ranking = -1, author = "George R. R. Martin", isbn = "978-0345540561"
		)
		val bookToCreate = persistBookRequest.toBook()
		val createdBook = bookToCreate.copy(id = "10")

		every { validationHandler.validate(Mono.just(persistBookRequest)) } returns Mono.just(persistBookRequest)
		every { bookRepository.save(bookToCreate) } returns Mono.just(createdBook)

		webClient.post().uri("")
			.accept(MediaType.APPLICATION_JSON)
			.contentType(MediaType.APPLICATION_JSON)
			.bodyValue(persistBookRequest)
			.exchange()
			.expectStatus().isCreated
			.expectBody().json(objectMapper.writeValueAsString(createdBook))

		verify(exactly = 1) { bookRepository.save(any<Book>()) }
		confirmVerified(bookRepository)
	}

	@Test
	fun `find book`() {
		val id = "1"
		val book = Book(id, "A Dance with Dragons", 1, "George R. R. Martin", "978-0345540560")

		every { bookRepository.findById(id) } returns Mono.just(book)

		webClient.get().uri("/$id")
			.accept(MediaType.APPLICATION_JSON)
			.header("Content-type", "application/json")
			.exchange()
			.expectStatus().isOk
			.expectBody().json(objectMapper.writeValueAsString(book))

		verify(exactly = 1) { bookRepository.findById(id) }
		confirmVerified(bookRepository)
	}

	@Test
	fun `find book by non existing id throws exception`() {
		val id = "99"

		every { bookRepository.findById(id) } returns Mono.empty()

		webClient.get().uri("/$id")
			.accept(MediaType.APPLICATION_JSON)
			.header("Content-type", "application/json")
			.exchange()
			.expectStatus().isNotFound

		verify(exactly = 1) { bookRepository.findById(id) }
		confirmVerified(bookRepository)
	}

	@Test
	fun `update book`() {
		val id = "1"
		val existingBook =
			Book(id, "A Dream of Spring", 0, "George R. R. Martin", "978-0345540561")
		val persistBookRequest =
			PersistBookRequest(id, "A Dream of Spring", 5, "George R. R. Martin", "978-0345540561")
		val updatedBook = existingBook.copy(ranking = 5)

		//every { entityValidator.validate(Mono.just(persistBookRequest)) } returns Mono.just(persistBookRequest)
		every { bookRepository.findById(id) } returns Mono.just(existingBook)
		every { bookRepository.save(updatedBook) } returns Mono.just(updatedBook)

		webClient.put().uri("/$id")
			.accept(MediaType.APPLICATION_JSON)
			.contentType(MediaType.APPLICATION_JSON)
			.bodyValue(persistBookRequest)
			.exchange()
			.expectStatus().isOk
			.expectBody().json(objectMapper.writeValueAsString(updatedBook))

		verify(exactly = 1) {
			bookRepository.findById(id)
			bookRepository.save(any<Book>())
		}
		confirmVerified(bookRepository)
	}

	@Test
	fun `update book by non existing id`() {
		val id = "1"
		val persistBookRequest =
			PersistBookRequest(id, "A Dream of Spring", 5, "George R. R. Martin", "978-0345540561")

		every { bookRepository.findById(id) } returns Mono.empty()

		webClient.put().uri("/$id")
			.accept(MediaType.APPLICATION_JSON)
			.contentType(MediaType.APPLICATION_JSON)
			.bodyValue(persistBookRequest)
			.exchange()
			.expectStatus().isNotFound

		verify(exactly = 1) { bookRepository.findById(id) }
		verify(exactly = 0) { bookRepository.save(any<Book>()) }
		confirmVerified(bookRepository)
	}

	@Test
	fun `delete book`() {
		val id = "1"
		val existingBook =
			Book(id, "A Dream of Spring", 5, "George R. R. Martin", "978-0345540561")

		every { bookRepository.findById(id) } returns Mono.just(existingBook)
		every { bookRepository.delete(existingBook) } returns Mono.empty()

		webClient.delete().uri("/$id")
			.accept(MediaType.APPLICATION_JSON)
			.header("Content-type", "application/json")
			.exchange()
			.expectStatus().isNoContent

		verify(exactly = 1) {
			bookRepository.findById(id)
			bookRepository.delete(existingBook)
		}
		confirmVerified(bookRepository)
	}

	@Test
	fun `delete book by non existing id throws exception`() {
		val id = "1"

		every { bookRepository.findById(id) } returns Mono.empty()

		webClient.delete().uri("/$id")
			.accept(MediaType.APPLICATION_JSON)
			.header("Content-type", "application/json")
			.exchange()
			.expectStatus().isNotFound

		verify(exactly = 1) { bookRepository.findById(id) }
		verify(exactly = 0) { bookRepository.delete(any()) }
		confirmVerified(bookRepository)
	}
}