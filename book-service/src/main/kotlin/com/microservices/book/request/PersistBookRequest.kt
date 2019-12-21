package com.microservices.book.request

import javax.validation.constraints.Max
import javax.validation.constraints.Min
import javax.validation.constraints.NotEmpty

data class PersistBookRequest(
	val id: String? = null,

	@field:NotEmpty(message = "Name is required")
	val name: String,

	@field:Min(0, message = "Ranking can't be negative")
	@field:Max(5, message = "Ranking must be less or equal than 5")
	val ranking: Int,

	@field:NotEmpty(message = "Author is required")
	val author: String,

	@field:NotEmpty(message = "ISBN is required")
	val isbn: String
)