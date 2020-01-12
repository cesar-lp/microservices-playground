package main

import (
	"strings"

	"github.com/cesar-lp/microservices-playground/movie-service/main/server"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func init() {
	viper.SetConfigFile("../config.json")

	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	env := viper.GetString("env")
	log.Info("Service running on " + strings.ToUpper(env))
}

/*
	TODO:
		Dockerfile
		Minor improvements (possible bug in update method)
		Cleanup
*/
func main() {
	env := viper.GetString("env")
	debug := viper.GetBool(env + ".debug")
	address := viper.GetString("server.address")

	dbHost := viper.GetString("db.host")
	dbPort := viper.GetInt("db.port")
	dbUser := viper.GetString("db.user")
	dbPassword := viper.GetString("db.password")
	dbName := viper.GetString("db.name")

	app := server.Configure(dbHost, dbPort, dbUser, dbPassword, dbName, debug)
	app.Run(address)
}
