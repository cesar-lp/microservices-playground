package repositories

import (
	"github.com/cesar-lp/microservices-playground/movie-service/main/database"
	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/jinzhu/gorm"
)

type MovieRepository interface {
	GetAllMovies() ([]models.Movie, error)
	GetMovieById(id int) (models.Movie, error)
	CreateMovie(movie *models.Movie) (models.Movie, error)
	UpdateMovie(id int, movie *models.Movie) (models.Movie, error)
	DeleteMovieById(id int) error
}

type MovieStore struct{}

func getDB() *gorm.DB {
	return database.Get().Debug().Model(&models.Movie{})
}

func (MovieStore) GetAllMovies() ([]models.Movie, error) {
	movies := []models.Movie{}
	var err = getDB().Limit(50).Find(&movies).Error
	return movies, err
}

func (MovieStore) GetMovieById(id int) (models.Movie, error) {
	var movie models.Movie
	var err = getDB().Where("id = ?", id).Take(&movie).Error
	return movie, err
}

func (MovieStore) CreateMovie(movie *models.Movie) (models.Movie, error) {
	return *movie, getDB().Create(&movie).Error
}

func (store MovieStore) UpdateMovie(id int, movieToUpdate *models.Movie) (models.Movie, error) {
	var err = getDB().Where("id = ?", id).Update(&movieToUpdate).Error
	return *movieToUpdate, err
}

func (store MovieStore) DeleteMovieById(id int) error {
	movie, err := store.GetMovieById(id)
	if err != nil {
		return err
	}
	return getDB().Delete(movie).Error
}
