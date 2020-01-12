package repositories_mocks

import (
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/models"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/repositories"

	"github.com/jinzhu/gorm"
)

type movieRepositoryMock struct {
	isEmpty bool
}

func MovieRepositoryMock(isEmpty bool) repositories.MovieRepository {
	return movieRepositoryMock{isEmpty: isEmpty}
}

func (repository movieRepositoryMock) GetAllMovies() ([]models.Movie, int64, error) {
	var movies []models.Movie

	if repository.isEmpty {
		return movies, int64(0), nil
	}

	movies = append(movies, models.Movie{Id: 1, Name: "Inception", Rating: 5})
	return movies, int64(1), nil
}

func (repository movieRepositoryMock) GetMovieById(id int) (models.Movie, int64, error) {
	if repository.isEmpty {
		return models.Movie{}, int64(0), gorm.ErrRecordNotFound
	}
	return models.Movie{Id: 1, Name: "Inception", Rating: 5}, int64(1), nil
}

func (movieRepositoryMock) CreateMovie(movie *models.Movie) (models.Movie, int64, error) {
	movie.Id = 1
	return *movie, int64(1), nil
}

func (repository movieRepositoryMock) UpdateMovie(id int, movie *models.Movie) (models.Movie, int64, error) {
	if repository.isEmpty {
		return models.Movie{}, int64(0), gorm.ErrRecordNotFound
	}
	return *movie, int64(1), nil
}

func (repository movieRepositoryMock) DeleteMovieById(id int) (int64, error) {
	if repository.isEmpty {
		return int64(0), gorm.ErrRecordNotFound
	}
	return int64(1), nil
}
