package repositories

import (
	"github.com/cesar-lp/microservices-playground/movie-service/main/models"

	"github.com/jinzhu/gorm"
)

// MovieRepository contract.
type MovieRepository interface {
	// GetAllMovies returns all movies, the number of modified rows and any error encountered.
	GetAllMovies() ([]models.Movie, int64, error)

	// GetMovieById returns a movie for a given id, the number of modified rows and any error encountered.
	GetMovieById(id int) (models.Movie, int64, error)

	// CreateMovie saves a movie to database and returns
	// the created movie, the number of modified rows and any error encountered.
	CreateMovie(movie *models.Movie) (models.Movie, int64, error)

	// UpdateMovie updates an existing and returns the updated movie, the number of modified rows
	// and any error encountered.
	UpdateMovie(id int, movie *models.Movie) (models.Movie, int64, error)

	// DeleteMovieById deletes a movie for a given id and returns the number of modified rows
	// and any error encountered.
	DeleteMovieById(id int) (int64, error)
}

type movieRepository struct {
	db *gorm.DB
}

// MovieRepositoryImpl builds and returns a movie repository implementation.
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
