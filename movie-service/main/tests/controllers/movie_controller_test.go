package controllers_tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cesar-lp/microservices-playground/movie-service/main/app/common"
	ctrl "github.com/cesar-lp/microservices-playground/movie-service/main/app/controllers"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/middlewares"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/models"
	mocks "github.com/cesar-lp/microservices-playground/movie-service/tests/handlers/mocks"
	"github.com/cesar-lp/microservices-playground/movie-service/tests/utils"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupController(isRepositoryEmpty bool) *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.JSONMiddleware)
	ctrl.MovieController(mocks.MovieHandlerMock(isRepositoryEmpty), r)
	return r
}

func setupTest(r *mux.Router, httpMethod string, uri string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(httpMethod, uri, body)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	return rr
}

func TestGetAllMovies(t *testing.T) {
	assert := assert.New(t)
	rr := setupTest(setupController(false), "GET", "/api/movies", nil)

	expectedResponseBody := [3]models.Movie{
		models.Movie{Id: 1, Name: "Inception", Rating: 5},
		models.Movie{Id: 2, Name: "The Dark Night", Rating: 4},
		models.Movie{Id: 3, Name: "Jumanji", Rating: 2},
	}

	var responseBody [3]models.Movie
	json.NewDecoder(rr.Body).Decode(&responseBody)

	assert.Equal(200, rr.Result().StatusCode)
	assert.Equal(expectedResponseBody, responseBody)
}

func TestGetAllMovies_ReturnsEmpty(t *testing.T) {
	assert := assert.New(t)
	rr := setupTest(setupController(true), "GET", "/api/movies", nil)

	var expectedResponseBody []models.Movie

	var responseBody []models.Movie
	json.NewDecoder(rr.Body).Decode(&responseBody)

	assert.Equal(200, rr.Result().StatusCode)
	assert.Equal(expectedResponseBody, responseBody)
}

func TestGetMovieById(t *testing.T) {
	assert := assert.New(t)
	rr := setupTest(setupController(false), "GET", "/api/movies/1", nil)

	expectedResponseBody := models.Movie{Id: 1, Name: "Inception", Rating: 5}

	var responseBody models.Movie
	json.NewDecoder(rr.Body).Decode(&responseBody)

	assert.Equal(200, rr.Result().StatusCode)
	assert.Equal(expectedResponseBody, responseBody)
}

func TestGetMovieById_ReturnsResourceNotFoundError(t *testing.T) {
	path := "/api/movies/1"
	assert := assert.New(t)
	rr := setupTest(setupController(true), "GET", path, nil)

	expectedResponseBody := ctrl.ResourceNotFound(path, "Movie not found for id 1")

	var responseBody ctrl.BaseErrorResponse
	json.NewDecoder(rr.Body).Decode(&responseBody)
	responseBody.Timestamp = expectedResponseBody.Timestamp

	assert.Equal(404, rr.Result().StatusCode)
	assert.Equal(expectedResponseBody, responseBody)
}

func TestCreateMovie(t *testing.T) {
	assert := assert.New(t)
	newMovie := models.Movie{Name: "A Quiet Place 2"}
	rr := setupTest(setupController(false), "POST", "/api/movies", utils.WriteBody(newMovie))

	expectedResponseBody := models.Movie{Id: 4, Name: "A Quiet Place 2", Rating: 0}

	var responseBody models.Movie
	json.NewDecoder(rr.Body).Decode(&responseBody)

	assert.Equal(201, rr.Result().StatusCode)
	assert.Equal(expectedResponseBody, responseBody)
}

func TestCreateMovie_ReturnsInvalidFieldsError(t *testing.T) {
	path := "/api/movies"
	assert := assert.New(t)
	rr := setupTest(setupController(false), "POST", path, utils.WriteBody(models.Movie{}))

	expectedResponseBody := ctrl.ValidationError(path, []common.FieldError{
		common.NewFieldError("name", "Name is required", ""),
	})

	var responseBody ctrl.ValidationErrorResponse
	json.NewDecoder(rr.Body).Decode(&responseBody)
	responseBody.Timestamp = expectedResponseBody.Timestamp

	assert.Equal(422, rr.Result().StatusCode)
	assert.Equal(expectedResponseBody, responseBody)
}

func TestUpdateMovie(t *testing.T) {
	path := "/api/movies/1"
	updatedMovie := models.Movie{Id: 1, Name: "Inception", Rating: 5}
	assert := assert.New(t)
	rr := setupTest(setupController(false), "PUT", path, utils.WriteBody(updatedMovie))

	expectedResponseBody := updatedMovie

	var responseBody models.Movie
	json.NewDecoder(rr.Body).Decode(&responseBody)

	assert.Equal(200, rr.Result().StatusCode)
	assert.Equal(expectedResponseBody, responseBody)
}

func TestUpdateMovie_ReturnsResourceNotFoundError(t *testing.T) {
	path := "/api/movies/1"
	assert := assert.New(t)
	updatedMovie := models.Movie{Id: 1, Name: "Inception", Rating: 5}
	rr := setupTest(setupController(true), "PUT", path, utils.WriteBody(updatedMovie))

	expectedResponseBody := ctrl.ResourceNotFound(path, "Movie not found for id 1")

	var responseBody ctrl.BaseErrorResponse
	json.NewDecoder(rr.Body).Decode(&responseBody)
	responseBody.Timestamp = expectedResponseBody.Timestamp

	assert.Equal(404, rr.Result().StatusCode)
	assert.Equal(expectedResponseBody, responseBody)
}

func TestDeleteMovieById(t *testing.T) {
	assert := assert.New(t)
	rr := setupTest(setupController(false), "DELETE", "/api/movies/1", nil)

	assert.Equal(204, rr.Result().StatusCode)
	assert.Nil(rr.Body.Bytes())
}

func TestDeleteMovieById_ReturnsResourceNotFoundError(t *testing.T) {
	path := "/api/movies/1"
	assert := assert.New(t)
	rr := setupTest(setupController(true), "DELETE", path, nil)

	expectedResponseBody := ctrl.ResourceNotFound(path, "Movie not found for id 1")

	var responseBody ctrl.BaseErrorResponse
	json.NewDecoder(rr.Body).Decode(&responseBody)
	responseBody.Timestamp = expectedResponseBody.Timestamp

	assert.Equal(404, rr.Result().StatusCode)
	assert.Equal(expectedResponseBody, responseBody)
}
