package repositories

import (
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

type movieRepository struct {
	db *gorm.DB
}

func MovieRepositoryImpl(database *gorm.DB) MovieRepository {
	return &movieRepository{
		db: database.Model(&models.Movie{}),
	}
}

func (r movieRepository) GetAllMovies() ([]models.Movie, int64, error) {
	movies := []models.Movie{}
	db := r.db.Find(&movies)
	return movies, db.RowsAffected, db.Error
}

func (r movieRepository) GetMovieById(id int) (models.Movie, int64, error) {
	var movie models.Movie
	db := r.db.Where("id = ?", id).Take(&movie)

	if db.RowsAffected == int64(0) {
		return movie, db.RowsAffected, gorm.ErrRecordNotFound
	}
	return movie, db.RowsAffected, db.Error
}

func (r movieRepository) CreateMovie(movie *models.Movie) (models.Movie, int64, error) {
	db := r.db.Create(&movie)
	return *movie, db.RowsAffected, db.Error
}

func (r movieRepository) UpdateMovie(id int, movieToUpdate *models.Movie) (models.Movie, int64, error) {
	db := r.db.Where("id = ?", id).Update(&movieToUpdate)

	if db.RowsAffected == int64(0) {
		return *movieToUpdate, db.RowsAffected, gorm.ErrRecordNotFound
	}
	return *movieToUpdate, db.RowsAffected, db.Error
}

func (r movieRepository) DeleteMovieById(id int) (int64, error) {
	db := r.db.Delete(models.Movie{Id: id})

	if db.RowsAffected == int64(0) {
		return db.RowsAffected, gorm.ErrRecordNotFound
	}
	return db.RowsAffected, db.Error
}
