package handlers

import (
	"errors"
	"strconv"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/common"
	. "github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/cesar-lp/microservices-playground/movie-service/main/repositories"
	"github.com/jinzhu/gorm"
)

type MovieHandlerAPI interface {
	GetAllMovies() HandlerResponse
	GetMovieById(id int) HandlerResponse
	CreateMovie(newMovie *Movie) HandlerResponse
	UpdateMovie(id int, movieToUpdate *Movie) HandlerResponse
	DeleteMovieById(id int) HandlerResponse
}

type movieHandler struct {
	repository repositories.MovieRepository
}

func CreateMovieHandler(movieRepository repositories.MovieRepository) MovieHandlerAPI {
	return &movieHandler{
		repository: movieRepository,
	}
}

// GetAll returns all movies
func (h *movieHandler) GetAllMovies() HandlerResponse {
	movies, _, err := h.repository.GetAllMovies()
	if err != nil {
		return InternalServerError(err)
	}
	return Ok(movies)
}

// Get returns a Movie for a given id
func (h *movieHandler) GetMovieById(id int) HandlerResponse {
	movie, _, err := h.repository.GetMovieById(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + strconv.Itoa(id)))
		}
		return InternalServerError(err)
	}
	return Ok(movie)
}

// Save a movie
func (h *movieHandler) CreateMovie(newMovie *Movie) HandlerResponse {
	fieldErrors := newMovie.Validate()

	if len(fieldErrors) > 0 {
		return UnprocessableEntity(fieldErrors)
	}

	newMovie.Initialize()
	createdMovie, _, err := h.repository.CreateMovie(newMovie)

	if err != nil {
		return InternalServerError(err)
	}
	return Created(createdMovie)
}

// Update a movie
func (h *movieHandler) UpdateMovie(id int, movieToUpdate *Movie) HandlerResponse {
	fieldErrors := movieToUpdate.Validate()

	if len(fieldErrors) > 0 {
		return UnprocessableEntity(fieldErrors)
	}

	updatedMovie, _, err := h.repository.UpdateMovie(id, movieToUpdate)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + strconv.Itoa(id)))
		}
		return InternalServerError(err)
	}
	return Ok(updatedMovie)
}

// Delete a movie for a given id
func (h *movieHandler) DeleteMovieById(id int) HandlerResponse {
	_, err := h.repository.DeleteMovieById(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + strconv.Itoa(id)))
		}
		return InternalServerError(err)
	}
	return NoContent()
}
