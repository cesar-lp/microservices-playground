package handlers_test

import (
	"strconv"
	"testing"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/common"
	"github.com/cesar-lp/microservices-playground/movie-service/main/handlers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	. "github.com/cesar-lp/microservices-playground/movie-service/main/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (*assert.Assertions, handlers.MovieHandlerAPI) {
	return assert.New(t), handlers.CreateMovieHandler(GetMovieStoreMock())
}

func TestGetAllMovies(t *testing.T) {
	assert, handler := setupTest(t)

	expectedMovies := []models.Movie{
		models.Movie{Id: 1, Name: "Inception", Rating: 5},
	}

	hr := handler.GetAllMovies()
	assert.Nil(hr.Err)
	assert.Nil(hr.FieldErrors)
	assert.Equal(expectedMovies, hr.Payload)
}

func TestGetMovieById(t *testing.T) {
	assert, handler := setupTest(t)

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
	assert, handler := setupTest(t)
	assertHandlerReturnsResourceNotFoundError(assert, handler.GetMovieById(id), id)
}

func TestCreateMovie(t *testing.T) {
	assert, handler := setupTest(t)

	movieToCreate := models.Movie{Name: "Inception", Rating: 5}
	expectedMovieCreated := models.Movie{Id: 1, Name: "Inception", Rating: 5}

	hr := handler.CreateMovie(&movieToCreate)
	assert.Nil(hr.Err)
	assert.Nil(hr.FieldErrors)
	assert.Equal(201, hr.StatusCode)
	assert.Equal(expectedMovieCreated, hr.Payload)
}

func TestCreateMovie_ReturnsInvalidFieldsError(t *testing.T) {
	assert, handler := setupTest(t)

	movieToCreate := models.Movie{}
	expectedFieldErrors := []FieldError{
		FieldError{FieldName: "name", Error: "Name is required", InvalidValue: ""},
	}

	hr := handler.CreateMovie(&movieToCreate)
	assert.Nil(hr.Err)
	assert.Nil(hr.Payload)
	assert.Equal(422, hr.StatusCode)
	assert.Equal(expectedFieldErrors, hr.FieldErrors)
}

func TestUpdateMovie(t *testing.T) {
	assert, handler := setupTest(t)

	id := 1
	movieToUpdate := models.Movie{Id: 1, Name: "Inception", Rating: 5}

	hr := handler.UpdateMovie(id, &movieToUpdate)
	assert.Nil(hr.Err)
	assert.Nil(hr.FieldErrors)
	assert.Equal(200, hr.StatusCode)
	assert.Equal(movieToUpdate, hr.Payload)
}

func TestUpdateMovie_ReturnsResourceNotFoundError(t *testing.T) {
	assert, handler := setupTest(t)

	id := -1
	movieToUpdate := models.Movie{Id: 1, Name: "Inception", Rating: 5}

	assertHandlerReturnsResourceNotFoundError(assert, handler.UpdateMovie(id, &movieToUpdate), id)
}

func TestDeleteMovieById(t *testing.T) {
	assert, handler := setupTest(t)

	id := 1

	hr := handler.DeleteMovieById(id)
	assert.Nil(hr.Payload)
	assert.Nil(hr.Err)
	assert.Nil(hr.FieldErrors)
	assert.Equal(204, hr.StatusCode)
}

func TestDeleteMovieById_ReturnsResourceNotFoundError(t *testing.T) {
	id := -1
	assert, handler := setupTest(t)
	assertHandlerReturnsResourceNotFoundError(assert, handler.DeleteMovieById(id), id)
}

func assertHandlerReturnsResourceNotFoundError(assert *assert.Assertions, hr HandlerResponse, id int) {
	assert.Nil(hr.Payload)
	assert.Nil(hr.FieldErrors)
	assert.Equal(404, hr.StatusCode)
	assert.Equal("Movie not found for id "+strconv.Itoa(id), hr.Err.Error())
}
