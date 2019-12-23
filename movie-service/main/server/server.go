package server

import (
	"fmt"
	"log"
	"net/http"

	ctrl "github.com/cesar-lp/microservices-playground/movie-service/main/controllers"
	mw "github.com/cesar-lp/microservices-playground/movie-service/main/middleware"
	"github.com/gorilla/mux"
)

// Server structure
type Server struct {
	Router *mux.Router
}

// Configure and returns a server instance
func Configure() Server {
	server := Server{}
	server.Router = getRoutes()
	return server
}

func getRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", ctrl.GetMovies).Methods("GET")
	r.HandleFunc("/api/movies", ctrl.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", ctrl.GetMovieByID).Methods("GET")
	r.HandleFunc("/api/movies/{id}", ctrl.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", ctrl.DeleteMovie).Methods("DELETE")

	r.Use(mw.JSONMiddleware)
	return r
}

// Run the server
func (server *Server) Run(addr string) {
	fmt.Println("Server up and running: listening to port 8081")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
