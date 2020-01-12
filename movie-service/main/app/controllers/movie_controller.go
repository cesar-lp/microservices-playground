package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cesar-lp/microservices-playground/movie-service/main/app/handlers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/models"

	"github.com/gorilla/mux"
)

type movieController struct {
	handler handlers.MovieHandler
}

// MovieController instantiates a movie controller.
func MovieController(h handlers.MovieHandler, r *mux.Router) {
	ctrl := movieController{
		handler: h,
	}

	r.HandleFunc("/api/movies", ctrl.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movies", ctrl.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", ctrl.GetMovieById).Methods("GET")
	r.HandleFunc("/api/movies/{id}", ctrl.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", ctrl.DeleteMovieById).Methods("DELETE")
}

// GetMovies returns all movies.
func (ctrl movieController) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	ServerResponse(w, r, ctrl.handler.GetAllMovies())
}

// GetMovieById returns a movie for a given id.
func (ctrl movieController) GetMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	ServerResponse(w, r, ctrl.handler.GetMovieById(id))
}

// CreateMovie creates and returns the created movie.
func (ctrl movieController) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	hr := ctrl.handler.CreateMovie(&movie)
	w.Header().Set("Location", r.RequestURI+"/"+strconv.Itoa(movie.Id))
	ServerResponse(w, r, hr)
}

// UpdateMovie updates a movie.
func (ctrl movieController) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedMovie models.Movie
	json.NewDecoder(r.Body).Decode(&updatedMovie)

	ServerResponse(w, r, ctrl.handler.UpdateMovie(id, &updatedMovie))
}

// DeleteMovieById deletes a movie for a given id.
func (ctrl movieController) DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	ServerResponse(w, r, ctrl.handler.DeleteMovieById(id))
}
