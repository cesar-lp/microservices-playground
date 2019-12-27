package main

import (
	db "github.com/cesar-lp/microservices-playground/movie-service/main/database"
	"github.com/cesar-lp/microservices-playground/movie-service/main/server"
	_ "github.com/lib/pq"
)

/*
	TODO:
		Logging
		Tests
		Retrieve values from environment
		Minor improvements
		Cleanup
*/
func main() {
	app := server.Configure()
	db.Load()
	app.Run("8081")
}
