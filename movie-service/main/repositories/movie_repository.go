package repositories

import (
	"github.com/cesar-lp/microservices-playground/movie-service/main/database"
	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/jinzhu/gorm"
)

type MovieRepository interface {
	GetAllMovies() ([]models.Movie, int64, error)
	GetMovieById(id int) (models.Movie, int64, error)
	CreateMovie(movie *models.Movie) (models.Movie, int64, error)
	UpdateMovie(id int, movie *models.Movie) (models.Movie, int64, error)
	DeleteMovieById(id int) (int64, error)
}

type MovieStore struct{}

func getDB() *gorm.DB {
	return database.Get().Debug().Model(&models.Movie{})
}

func (MovieStore) GetAllMovies() ([]models.Movie, int64, error) {
	movies := []models.Movie{}
	db := getDB().Limit(50).Find(&movies)
	return movies, db.RowsAffected, db.Error
}

func (MovieStore) GetMovieById(id int) (models.Movie, int64, error) {
	var movie models.Movie
	db := getDB().Where("id = ?", id).Take(&movie)

	if db.RowsAffected == int64(0) {
		return movie, db.RowsAffected, gorm.ErrRecordNotFound
	}
	return movie, db.RowsAffected, db.Error
}

func (MovieStore) CreateMovie(movie *models.Movie) (models.Movie, int64, error) {
	db := getDB().Create(&movie)
	return *movie, db.RowsAffected, db.Error
}

func (store MovieStore) UpdateMovie(id int, movieToUpdate *models.Movie) (models.Movie, int64, error) {
	db := getDB().Where("id = ?", id).Update(&movieToUpdate)

	if db.RowsAffected == int64(0) {
		return *movieToUpdate, db.RowsAffected, gorm.ErrRecordNotFound
	}
	return *movieToUpdate, db.RowsAffected, db.Error
}

func (store MovieStore) DeleteMovieById(id int) (int64, error) {
	movie := models.Movie{Id: id}
	db := getDB().Delete(&movie)

	if db.RowsAffected == int64(0) {
		return db.RowsAffected, gorm.ErrRecordNotFound
	}
	return db.RowsAffected, db.Error
}
