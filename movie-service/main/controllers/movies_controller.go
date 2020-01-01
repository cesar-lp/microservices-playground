package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/common"
	handler "github.com/cesar-lp/microservices-playground/movie-service/main/handlers"
	. "github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/gorilla/mux"
)

// GetMovies returns all available movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	ServerResponse(w, r, handler.GetAllMovies())
}

// GetMovieByID returns a Movie for a given id
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	ServerResponse(w, r, handler.GetMovieById(id))
}

// CreateMovie creates and returns the created movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)

	hr := handler.CreateMovie(&movie)
	w.Header().Set("Location", r.RequestURI+"/"+strconv.Itoa(movie.Id))
	ServerResponse(w, r, hr)
}

// UpdateMovie updates a movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedMovie Movie
	json.NewDecoder(r.Body).Decode(&updatedMovie)

	ServerResponse(w, r, handler.UpdateMovie(id, &updatedMovie))
}

// DeleteMovie deletes a movie for a given ID
func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	ServerResponse(w, r, handler.DeleteMovieById(id))
}
