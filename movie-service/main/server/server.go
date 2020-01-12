package server

import (
	"net/http"

	"github.com/cesar-lp/microservices-playground/movie-service/main/controllers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/handlers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/middlewares"
	"github.com/cesar-lp/microservices-playground/movie-service/main/repositories"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server structure.
type Server struct {
	db     Database
	router *mux.Router
}

// Configure a server instance.
func Configure(host string, port int, user, password, name string, logDB bool) Server {
	server := Server{
		db:     setupDB(host, port, user, password, name, logDB),
		router: setupRouter(),
	}

	return server
}

// Run the server.
func (server *Server) Run(address string) {
	server.db.Connect()
	server.db.LoadSeeds()
	server.initializeDependencies()
	defer server.db.instance.Close()

	log.Info("Server up and running, listening to port" + address)
	log.Panic(http.ListenAndServe(address, server.router))
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
