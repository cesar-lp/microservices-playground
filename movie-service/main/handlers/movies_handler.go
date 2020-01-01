package handlers

import (
	"errors"
	"strconv"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/common"
	. "github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/cesar-lp/microservices-playground/movie-service/main/repositories"
	"github.com/jinzhu/gorm"
)

var movieRepository = repositories.MovieStore{}

// GetAll returns all movies
func GetAllMovies() HandlerResponse {
	movies, _, err := movieRepository.GetAllMovies()
	if err != nil {
		return InternalServerError(err)
	}
	return Ok(movies)
}

// Get returns a Movie for a given id
func GetMovieById(id int) HandlerResponse {
	movie, _, err := movieRepository.GetMovieById(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + strconv.Itoa(id)))
		}
		return InternalServerError(err)
	}
	return Ok(movie)
}

// Save a movie
func CreateMovie(newMovie *Movie) HandlerResponse {
	fieldErrors := newMovie.Validate()

	if len(fieldErrors) > 0 {
		return UnprocessableEntity(fieldErrors)
	}

	newMovie.Initialize()
	createdMovie, _, err := movieRepository.CreateMovie(newMovie)

	if err != nil {
		return InternalServerError(err)
	}
	return Created(createdMovie)
}

// Update a movie
func UpdateMovie(id int, movieToUpdate *Movie) HandlerResponse {
	fieldErrors := movieToUpdate.Validate()

	if len(fieldErrors) > 0 {
		return UnprocessableEntity(fieldErrors)
	}

	updatedMovie, _, err := movieRepository.UpdateMovie(id, movieToUpdate)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + strconv.Itoa(id)))
		}
		return InternalServerError(err)
	}
	return Ok(updatedMovie)
}

// Delete a movie for a given id
func DeleteMovieById(id int) HandlerResponse {
	_, err := movieRepository.DeleteMovieById(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + strconv.Itoa(id)))
		}
		return InternalServerError(err)
	}
	return NoContent()
}
