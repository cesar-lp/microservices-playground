package handlers

import (
	"errors"
	"strconv"

	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/cesar-lp/microservices-playground/movie-service/main/repositories"

	"github.com/jinzhu/gorm"
)

// MovieHandler contract.
type MovieHandler interface {
	// GetAllMovies returns a handler response containing all found movies
	// and any additional information.
	GetAllMovies() HandlerResponse

	// GetMovieById returns a handler response containing the movie found for a given id
	// and any additional information.
	GetMovieById(id int) HandlerResponse

	// CreateMovie creates a movie and returns a handler response containing the created movie
	// and any additional information.
	CreateMovie(newMovie *models.Movie) HandlerResponse

	// UpdateMovie updates a movie and returns a handler response containing the updated movie
	// and any additional information.
	UpdateMovie(id int, movieToUpdate *models.Movie) HandlerResponse

	// DeleteMovieById deletes a movie for a given id and returns a handler response containing
	// the results of the operation.
	DeleteMovieById(id int) HandlerResponse
}

type movieHandler struct {
	repository repositories.MovieRepository
}

// MovieHandlerImpl builds and returns a movie handler implementation.
func MovieHandlerImpl(r repositories.MovieRepository) MovieHandler {
	return &movieHandler{
		repository: r,
	}
}

func (h movieHandler) GetAllMovies() HandlerResponse {
	movies, _, err := h.repository.GetAllMovies()
	if err != nil {
		return InternalServerError(err)
	}
	return Ok(movies)
}

func (h movieHandler) GetMovieById(id int) HandlerResponse {
	movie, _, err := h.repository.GetMovieById(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + strconv.Itoa(id)))
		}
		return InternalServerError(err)
	}
	return Ok(movie)
}

func (h movieHandler) CreateMovie(newMovie *models.Movie) HandlerResponse {
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

func (h movieHandler) UpdateMovie(id int, movieToUpdate *models.Movie) HandlerResponse {
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

func (h movieHandler) DeleteMovieById(id int) HandlerResponse {
	_, err := h.repository.DeleteMovieById(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + strconv.Itoa(id)))
		}
		return InternalServerError(err)
	}
	return NoContent()
}
