package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/common"
	"github.com/cesar-lp/microservices-playground/movie-service/main/handlers"
	. "github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/gorilla/mux"
)

type MovieController struct {
	Handler handlers.MovieHandlerAPI
}

func CreateMovieController(h handlers.MovieHandlerAPI, r *mux.Router) {
	ctrl := &MovieController{
		Handler: h,
	}

	r.HandleFunc("/api/movies", ctrl.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movies", ctrl.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", ctrl.GetMovieById).Methods("GET")
	r.HandleFunc("/api/movies/{id}", ctrl.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", ctrl.DeleteMovieById).Methods("DELETE")
}

// GetMovies returns all available movies
func (ctrl MovieController) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	ServerResponse(w, r, ctrl.Handler.GetAllMovies())
}

// GetMovieByID returns a Movie for a given id
func (ctrl MovieController) GetMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	ServerResponse(w, r, ctrl.Handler.GetMovieById(id))
}

// CreateMovie creates and returns the created movie
func (ctrl MovieController) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)

	hr := ctrl.Handler.CreateMovie(&movie)
	w.Header().Set("Location", r.RequestURI+"/"+strconv.Itoa(movie.Id))
	ServerResponse(w, r, hr)
}

// UpdateMovie updates a movie
func (ctrl MovieController) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedMovie Movie
	json.NewDecoder(r.Body).Decode(&updatedMovie)

	ServerResponse(w, r, ctrl.Handler.UpdateMovie(id, &updatedMovie))
}

// DeleteMovie deletes a movie for a given ID
func (ctrl MovieController) DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	ServerResponse(w, r, ctrl.Handler.DeleteMovieById(id))
}
