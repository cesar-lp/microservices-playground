package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cesar-lp/microservices-playground/movie-service/main/controllers"
	db "github.com/cesar-lp/microservices-playground/movie-service/main/database"
	"github.com/cesar-lp/microservices-playground/movie-service/main/handlers"
	mw "github.com/cesar-lp/microservices-playground/movie-service/main/middleware"
	"github.com/cesar-lp/microservices-playground/movie-service/main/repositories"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server structure
type Server struct {
	Database *gorm.DB
	Router   *mux.Router
}

// Configure and returns a server instance
func Configure() Server {
	return Server{
		Database: db.Connect(),
		Router:   getRoutes(),
	}
}

func getRoutes() *mux.Router {
	r := mux.NewRouter()

	movieHandler := handlers.CreateMovieHandler(repositories.GetMovieRepository())
	controllers.CreateMovieController(movieHandler, r)

	r.Use(mw.JSONMiddleware)
	return r
}

// Run the server
func (server *Server) Run(port string) {
	fmt.Println("Server up and running: listening to port " + port)
	log.Fatal(http.ListenAndServe(":"+port, server.Router))
}
