package handlers_mocks

import (
	"errors"
	"strconv"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/common"
	"github.com/cesar-lp/microservices-playground/movie-service/main/handlers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/cesar-lp/microservices-playground/movie-service/main/repositories"
	repositories_mocks "github.com/cesar-lp/microservices-playground/movie-service/main/repositories/mocks"
)

// TODO:
// mock 422

type movieHandlerMock struct {
	repository        repositories.MovieRepository
	isRepositoryEmpty bool
}

func MovieHandlerMock(isRepositoryEmpty bool) handlers.MovieHandlerAPI {
	return movieHandlerMock{
		repository:        repositories_mocks.GetMovieStoreMock(),
		isRepositoryEmpty: isRepositoryEmpty,
	}
}

func (mock movieHandlerMock) GetAllMovies() HandlerResponse {
	var movies []models.Movie

	if mock.isRepositoryEmpty {
		return Ok(movies)
	}

	movies = append(movies,
		models.Movie{Id: 1, Name: "Inception", Rating: 5},
		models.Movie{Id: 2, Name: "The Dark Night", Rating: 4},
		models.Movie{Id: 3, Name: "Jumanji", Rating: 2},
	)
	return Ok(movies)
}

func (mock movieHandlerMock) GetMovieById(id int) HandlerResponse {
	if mock.isRepositoryEmpty {
		return getMovieNotFoundError(id)
	}
	return Ok(models.Movie{Id: 1, Name: "Inception", Rating: 5})
}

func (movieHandlerMock) CreateMovie(newMovie *models.Movie) HandlerResponse {
	if fieldErrors := newMovie.Validate(); len(fieldErrors) > 0 {
		return UnprocessableEntity(fieldErrors)
	}

	newMovie.Initialize()

	return Created(models.Movie{Id: 4, Name: "A Quiet Place 2", Rating: 0})
}

func (mock movieHandlerMock) UpdateMovie(id int, movieToUpdate *models.Movie) HandlerResponse {
	if mock.isRepositoryEmpty {
		return getMovieNotFoundError(id)
	}

	if fieldErrors := movieToUpdate.Validate(); len(fieldErrors) > 0 {
		return UnprocessableEntity(fieldErrors)
	}

	return Ok(*movieToUpdate)
}

func (mock movieHandlerMock) DeleteMovieById(id int) HandlerResponse {
	if mock.isRepositoryEmpty {
		return getMovieNotFoundError(id)
	}
	return NoContent()
}

func getMovieNotFoundError(id int) HandlerResponse {
	return NotFound(errors.New("Movie not found for id " + strconv.Itoa(id)))
}
