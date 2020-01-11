package main

import (
	"fmt"
	"strings"

	"github.com/cesar-lp/microservices-playground/movie-service/main/server"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func init() {
	viper.SetConfigFile("../config.json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	env := viper.GetString("env")
	fmt.Println("Service running on " + strings.ToUpper(env))
}

/*
	TODO:
		Migrate to Gin?
		Logging
		Dockerfile
		Minor improvements - Cleanup
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
