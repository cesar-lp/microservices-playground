package controllers

import (
	"encoding/json"
	"net/http"

	handler "github.com/cesar-lp/microservices-playground/movie-service/main/handlers"
	. "github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/gorilla/mux"
)

// GetMovies returns all available movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(handler.GetAll())
}

// GetMovieByID returns a Movie for a given ID
func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(handler.Get(params["id"]))
}

// CreateMovie creates and returns the created movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(handler.Save(movie))
}

// UpdateMovie updates a movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

	json.NewEncoder(w).Encode(handler.Update(params["id"], updatedMovie))
}

// DeleteMovie deletes a movie for a given ID
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	handler.Delete(params["id"])

	w.WriteHeader(http.StatusNoContent)
}
