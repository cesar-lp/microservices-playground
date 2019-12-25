package handlers

import (
	"errors"
	"math/rand"
	"strconv"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/common"
	. "github.com/cesar-lp/microservices-playground/movie-service/main/models"
)

var movies = []Movie{
	Movie{ID: "1", Name: "Inception", Rating: 5},
	Movie{ID: "2", Name: "Interstellar", Rating: 5},
	Movie{ID: "3", Name: "The Dark Knight", Rating: 5},
}

// GetAll returns all movies
func GetAll() HandlerResponse {
	return Ok(movies)
}

// Get returns a Movie for a given ID
func Get(id string) HandlerResponse {
	pos := -1

	for i, movie := range movies {
		if movie.ID == id {
			pos = i
			break
		}
	}

	if pos == -1 {
		return NotFound(errors.New("Movie not found for id " + id))
	}

	return Ok(movies[pos])
}

// Save a movie
func Save(newMovie Movie) HandlerResponse {
	fieldErrors := newMovie.Validate()

	if len(fieldErrors) > 0 {
		return UnprocessableEntity(fieldErrors)
	}

	newMovie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, newMovie)
	return Created(newMovie)
}

// Update a movie
func Update(id string, updatedMovie Movie) HandlerResponse {
	pos := -1
	fieldErrors := updatedMovie.Validate()

	if len(fieldErrors) > 0 {
		return UnprocessableEntity(fieldErrors)
	}

	for i, movie := range movies {
		if movie.ID == id {
			pos = i
			movies[i] = updatedMovie
			break
		}
	}

	if pos == -1 {
		return NotFound(errors.New("Movie not found for id " + id))
	}

	return Ok(movies[pos])
}

// Delete a movie for a given ID
func Delete(id string) HandlerResponse {
	notFound := true

	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			notFound = false
			break
		}
	}

	if notFound {
		return NotFound(errors.New("Movie not found for id " + id))
	}

	return NoContent()
}
