package repositories_mocks

import (
	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/cesar-lp/microservices-playground/movie-service/main/repositories"
	"github.com/jinzhu/gorm"
)

type movieStoreMock struct{}

func GetMovieStoreMock() repositories.MovieRepository {
	return movieStoreMock{}
}

func (mock movieStoreMock) GetAllMovies() ([]models.Movie, int64, error) {
	movies := []models.Movie{
		models.Movie{Id: 1, Name: "Inception", Rating: 5},
	}
	return movies, int64(1), nil
}

func (mock movieStoreMock) GetMovieById(id int) (models.Movie, int64, error) {
	if id == -1 {
		return models.Movie{}, int64(0), gorm.ErrRecordNotFound
	}
	return models.Movie{Id: 1, Name: "Inception", Rating: 5}, int64(1), nil
}

func (mock movieStoreMock) CreateMovie(movie *models.Movie) (models.Movie, int64, error) {
	movie.Id = 1
	return *movie, int64(1), nil
}

func (mock movieStoreMock) UpdateMovie(id int, movie *models.Movie) (models.Movie, int64, error) {
	if id == -1 {
		return models.Movie{}, int64(0), gorm.ErrRecordNotFound
	}
	return *movie, int64(1), nil
}

func (mock movieStoreMock) DeleteMovieById(id int) (int64, error) {
	if id == -1 {
		return int64(0), gorm.ErrRecordNotFound
	}
	return int64(1), nil
}
