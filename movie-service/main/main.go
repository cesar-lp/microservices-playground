package main

import (
	"strings"

	"github.com/cesar-lp/microservices-playground/movie-service/main/config"
	"github.com/cesar-lp/microservices-playground/movie-service/main/server"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.Load()

	log.Info("Service running on " + strings.ToUpper(config.Server.Environment))
	debug := viper.GetBool("SERVER_DB_LOG")

	app := server.Configure(config.DB, debug)
	app.Run(config.Server.Address)
}
