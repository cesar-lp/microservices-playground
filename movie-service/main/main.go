package main

import (
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/config"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/server"

	_ "github.com/lib/pq"
)

func main() {
	config := config.Load()
	app := server.Configure(config.DB)
	app.Run(config.Server)
}
