package com.microservices.book.entity

import org.springframework.data.annotation.Id
import org.springframework.data.mongodb.core.mapping.Document

@Document
data class Book(
	@Id
	val id: String? = null,
	val name: String,
	var ranking: Int,
	val author: String,
	val isbn: String
)