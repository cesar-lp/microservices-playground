package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/cesar-lp/microservices-playground/movie-service/main/app/config"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/controllers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/handlers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/middlewares"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/repositories"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server structure.
type Server struct {
	db     Database
	router *mux.Router
}

// Configure a server instance.
func Configure(dbConfig config.DBConfig) Server {
	server := Server{
		db:     setupDB(dbConfig),
		router: setupRouter(),
	}

	return server
}

// Run the server.
func (server *Server) Run(config config.ServerConfig) {
	server.db.Connect()
	server.db.LoadSeeds()
	server.initializeDependencies()
	defer server.db.instance.Close()

	log.Infof("Server up and running on %s, listening to port %d",
		strings.ToUpper(config.Environment), config.Port)
	log.Panic(http.ListenAndServe(":"+strconv.Itoa(config.Port), server.router))
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.JSONMiddleware)
	r.Use(middlewares.LoggingMiddleware)

	return r
}

func (server Server) initializeDependencies() {
	movieRepository := repositories.MovieRepositoryImpl(server.db.instance)
	movieHandler := handlers.MovieHandlerImpl(movieRepository)
	controllers.MovieController(movieHandler, server.router)
}
