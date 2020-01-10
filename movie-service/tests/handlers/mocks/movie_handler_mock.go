package handlers_mocks

import (
	"errors"
	"strconv"

	"github.com/cesar-lp/microservices-playground/movie-service/main/handlers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/cesar-lp/microservices-playground/movie-service/main/repositories"
	rMocks "github.com/cesar-lp/microservices-playground/movie-service/tests/repositories/mocks"
)

type movieHandlerMock struct {
	repository         repositories.MovieRepository
	hasEmptyRepository bool
}

func MovieHandlerMock(isRepositoryEmpty bool) handlers.MovieHandler {
	return movieHandlerMock{
		repository:         rMocks.MovieRepositoryMock(isRepositoryEmpty),
		hasEmptyRepository: isRepositoryEmpty,
	}
}

func (handler movieHandlerMock) GetAllMovies() handlers.HandlerResponse {
	var movies []models.Movie

	if handler.hasEmptyRepository {
		return handlers.Ok(movies)
	}

	movies = append(movies,
		models.Movie{Id: 1, Name: "Inception", Rating: 5},
		models.Movie{Id: 2, Name: "The Dark Night", Rating: 4},
		models.Movie{Id: 3, Name: "Jumanji", Rating: 2},
	)
	return handlers.Ok(movies)
}

func (handler movieHandlerMock) GetMovieById(id int) handlers.HandlerResponse {
	if handler.hasEmptyRepository {
		return getMovieNotFoundError(id)
	}
	return handlers.Ok(models.Movie{Id: 1, Name: "Inception", Rating: 5})
}

func (movieHandlerMock) CreateMovie(newMovie *models.Movie) handlers.HandlerResponse {
	if fieldErrors := newMovie.Validate(); len(fieldErrors) > 0 {
		return handlers.UnprocessableEntity(fieldErrors)
	}

	newMovie.Initialize()

	return handlers.Created(models.Movie{Id: 4, Name: "A Quiet Place 2", Rating: 0})
}

func (handler movieHandlerMock) UpdateMovie(id int, movieToUpdate *models.Movie) handlers.HandlerResponse {
	if handler.hasEmptyRepository {
		return getMovieNotFoundError(id)
	}

	if fieldErrors := movieToUpdate.Validate(); len(fieldErrors) > 0 {
		return handlers.UnprocessableEntity(fieldErrors)
	}

	return handlers.Ok(*movieToUpdate)
}

func (handler movieHandlerMock) DeleteMovieById(id int) handlers.HandlerResponse {
	if handler.hasEmptyRepository {
		return getMovieNotFoundError(id)
	}
	return handlers.NoContent()
}

func getMovieNotFoundError(id int) handlers.HandlerResponse {
	return handlers.NotFound(errors.New("Movie not found for id " + strconv.Itoa(id)))
}
