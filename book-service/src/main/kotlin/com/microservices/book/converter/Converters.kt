package com.microservices.book.converter

import com.microservices.book.entity.Book
import com.microservices.book.request.PersistBookRequest

fun PersistBookRequest.toBook() =
	Book(this.id, this.name, this.ranking, this.author, this.isbn)
