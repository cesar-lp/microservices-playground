package com.microservices.book.repository

import com.microservices.book.entity.Book
import org.springframework.data.mongodb.repository.ReactiveMongoRepository
import org.springframework.stereotype.Repository

@Repository
interface BookRepository : ReactiveMongoRepository<Book, String>