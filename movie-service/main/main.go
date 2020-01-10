package main

import (
	"github.com/cesar-lp/microservices-playground/movie-service/main/server"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "admin"
	dbPassword = "admin"
	dbName     = "movie-service"
)

/*
	TODO:
		Logging
		Retrieve values from environment
		Minor improvements
		Migrate to Gin?
		Cleanup
*/
func main() {
	logDB := false

	app := server.Configure(dbHost, dbPort, dbUser, dbPassword, dbName, logDB)
	app.Run("8081")
}
