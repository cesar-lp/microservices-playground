package handlers_test

import (
	"strconv"
	"testing"

	"github.com/cesar-lp/microservices-playground/movie-service/main/app/common"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/handlers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/models"
	mocks "github.com/cesar-lp/microservices-playground/movie-service/tests/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T, hasEmptyRepository bool) (*assert.Assertions, handlers.MovieHandler) {
	return assert.New(t), handlers.MovieHandlerImpl(mocks.MovieRepositoryMock(hasEmptyRepository))
}

func TestGetAllMovies(t *testing.T) {
	assert, handler := setupTest(t, false)

	expectedMovies := []models.Movie{
		models.Movie{Id: 1, Name: "Inception", Rating: 5},
	}

	hr := handler.GetAllMovies()
	assert.Nil(hr.Err)
	assert.Nil(hr.FieldErrors)
	assert.Equal(expectedMovies, hr.Payload)

	assert, handler = setupTest(t, true)
}

func TestGetAllMovies_ReturnsEmpty(t *testing.T) {
	assert, handler := setupTest(t, true)

	var expectedMovies []models.Movie

	hr := handler.GetAllMovies()
	assert.Nil(hr.Err)
	assert.Nil(hr.FieldErrors)
	assert.Equal(expectedMovies, hr.Payload)
}

func TestGetMovieById(t *testing.T) {
	assert, handler := setupTest(t, false)

	id := 1
	expectedMovie := models.Movie{Id: id, Name: "Inception", Rating: 5}

	hr := handler.GetMovieById(id)
	assert.Nil(hr.Err)
	assert.Nil(hr.FieldErrors)
	assert.Equal(200, hr.StatusCode)
	assert.Equal(expectedMovie, hr.Payload)
}

func TestGetMovieById_ReturnsResourceNotFoundError(t *testing.T) {
	id := -1
	assert, handler := setupTest(t, true)
	assertHandlerReturnsResourceNotFoundError(assert, handler.GetMovieById(id), id)
}

func TestCreateMovie(t *testing.T) {
	assert, handler := setupTest(t, false)

	movieToCreate := models.Movie{Name: "Inception", Rating: 5}
	expectedMovieCreated := models.Movie{Id: 1, Name: "Inception", Rating: 5}

	hr := handler.CreateMovie(&movieToCreate)
	assert.Nil(hr.Err)
	assert.Nil(hr.FieldErrors)
	assert.Equal(201, hr.StatusCode)
	assert.Equal(expectedMovieCreated, hr.Payload)
}

func TestCreateMovie_ReturnsInvalidFieldsError(t *testing.T) {
	assert, handler := setupTest(t, false)

	movieToCreate := models.Movie{}
	expectedFieldErrors := []common.FieldError{
		common.NewFieldError("name", "Name is required", ""),
	}

	hr := handler.CreateMovie(&movieToCreate)
	assert.Nil(hr.Err)
	assert.Nil(hr.Payload)
	assert.Equal(422, hr.StatusCode)
	assert.Equal(expectedFieldErrors, hr.FieldErrors)
}

func TestUpdateMovie(t *testing.T) {
	assert, handler := setupTest(t, false)

	id := 1
	movieToUpdate := models.Movie{Id: 1, Name: "Inception", Rating: 5}

	hr := handler.UpdateMovie(id, &movieToUpdate)
	assert.Nil(hr.Err)
	assert.Nil(hr.FieldErrors)
	assert.Equal(200, hr.StatusCode)
	assert.Equal(movieToUpdate, hr.Payload)
}

func TestUpdateMovie_ReturnsResourceNotFoundError(t *testing.T) {
	assert, handler := setupTest(t, true)

	id := -1
	movieToUpdate := models.Movie{Id: 1, Name: "Inception", Rating: 5}

	assertHandlerReturnsResourceNotFoundError(assert, handler.UpdateMovie(id, &movieToUpdate), id)
}

func TestDeleteMovieById(t *testing.T) {
	assert, handler := setupTest(t, false)

	id := 1

	hr := handler.DeleteMovieById(id)
	assert.Nil(hr.Payload)
	assert.Nil(hr.Err)
	assert.Nil(hr.FieldErrors)
	assert.Equal(204, hr.StatusCode)
}

func TestDeleteMovieById_ReturnsResourceNotFoundError(t *testing.T) {
	id := -1
	assert, handler := setupTest(t, true)
	assertHandlerReturnsResourceNotFoundError(assert, handler.DeleteMovieById(id), id)
}

func assertHandlerReturnsResourceNotFoundError(assert *assert.Assertions, hr handlers.HandlerResponse, id int) {
	assert.Nil(hr.Payload)
	assert.Nil(hr.FieldErrors)
	assert.Equal(404, hr.StatusCode)
	assert.Equal("Movie not found for id "+strconv.Itoa(id), hr.Err.Error())
}
