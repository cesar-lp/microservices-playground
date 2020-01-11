package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cesar-lp/microservices-playground/movie-service/main/controllers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/handlers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/middlewares"
	"github.com/cesar-lp/microservices-playground/movie-service/main/repositories"

	"github.com/gorilla/mux"
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
func (server *Server) Run(port string) {
	server.db.Connect()
	server.db.LoadSeeds()
	server.initializeDependencies()

	fmt.Println("Server up and running: listening to port " + port)
	log.Fatal(http.ListenAndServe(":"+port, server.router))
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.JSONMiddleware)
	return r
}

func (server Server) initializeDependencies() {
	movieRepository := repositories.MovieRepositoryImpl(server.db.instance)
	movieHandler := handlers.MovieHandlerImpl(movieRepository)
	controllers.MovieController(movieHandler, server.router)
}
