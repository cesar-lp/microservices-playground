package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/gorilla/mux"
)

// GetMovies returns all available movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.GetAllMovies())
}

// GetMovieByID returns a Movie for a given ID
func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(models.GetMovieByID(params["id"]))
}

// CreateMovie creates and returns the created movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateMovie(movie))
}

// UpdateMovie updates a movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedMovie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

	json.NewEncoder(w).Encode(models.UpdateMovie(params["id"], updatedMovie))
}

// DeleteMovie deletes a movie for a given ID
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	models.DeleteMovieByID(params["id"])
	w.WriteHeader(http.StatusNoContent)
}
