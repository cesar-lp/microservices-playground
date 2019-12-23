package main

import (
	"github.com/cesar-lp/microservices-playground/movie-service/main/server"
)

/*
	TODO:
		Retrieve from DB
		Logging
		Custom responses
		Exception handling
		Validation
		Tests
		Minor improvements
		Cleanup
*/

func main() {
	app := server.Configure()
	app.Run(":8081")
}
