package controllers

import (
	"encoding/json"
	"net/http"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/common"
	handler "github.com/cesar-lp/microservices-playground/movie-service/main/handlers"
	. "github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/gorilla/mux"
)

// GetMovies returns all available movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	ServerResponse(w, r, handler.GetAll())
}

// GetMovieByID returns a Movie for a given ID
func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ServerResponse(w, r, handler.Get(params["id"]))
}

// CreateMovie creates and returns the created movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	ServerResponse(w, r, handler.Save(movie))
}

// UpdateMovie updates a movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

	ServerResponse(w, r, handler.Update(params["id"], updatedMovie))
}

// DeleteMovie deletes a movie for a given ID
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ServerResponse(w, r, handler.Delete(params["id"]))
}
