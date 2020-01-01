package repositories

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cesar-lp/microservices-playground/movie-service/main/database"
	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

// TODO:
// remove global variables
// see if there's a better way to run before & after functions
// test create & update with invalid values

var db *sql.DB
var _gorm *gorm.DB
var mock sqlmock.Sqlmock

func setup() {
	_db, _mock, err := sqlmock.New()

	if err != nil {
		log.Fatalf("Error '%s' when opening a stub database connection", err)
	}
	db = _db

	gorm, err := gorm.Open("sqlmock", db)

	if err != nil {
		log.Fatal("Could not mock a gorm connection")
	}

	database.Configure(gorm)
	_gorm = gorm
	mock = _mock
}

func teardown() {
	db.Close()
	_gorm.Close()
}

func TestGetAllMovies(t *testing.T) {
	setup()
	store := MovieStore{}
	assert := assert.New(t)
	query := "^SELECT (.+) FROM \"movies\" LIMIT 50$"

	// no movies in store

	expectedRowsAffected := int64(0)

	mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{}))

	foundMovies, rowsAffected, err := store.GetAllMovies()

	assert.Nil(err)
	assert.Len(foundMovies, 0)
	assert.Equal(expectedRowsAffected, rowsAffected)

	// movies in store

	expectedRowsAffected = int64(3)
	expectedMovies := []models.Movie{
		models.Movie{Name: "Inception", Rating: 5},
		models.Movie{Name: "Interstellar", Rating: 5},
		models.Movie{Name: "The Dark Knight", Rating: 5},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "rating"})

	for _, m := range expectedMovies {
		rows.AddRow(m.Id, m.Name, m.Rating)
	}

	mock.ExpectQuery(query).WillReturnRows(rows)

	foundMovies, rowsAffected, err = store.GetAllMovies()

	assert.Nil(err)
	assert.Len(foundMovies, len(expectedMovies))
	assert.Equal(expectedMovies, foundMovies)
	assert.Equal(expectedRowsAffected, rowsAffected)
	teardown()
}

func TestGetMovieById(t *testing.T) {
	setup()
	store := MovieStore{}
	assert := assert.New(t)
	query := "^SELECT (.+) FROM \"movies\" WHERE \\(id = \\?\\) LIMIT 1$"

	expectedRowsAffected := int64(1)
	expectedMovie := models.Movie{
		Id:     1,
		Name:   "Inception",
		Rating: 5,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "rating"}).
		AddRow(expectedMovie.Id, expectedMovie.Name, expectedMovie.Rating)

	mock.ExpectQuery(query).
		WithArgs(1).
		WillReturnRows(rows)

	foundMovie, rowsAffected, err := store.GetMovieById(1)

	assert.Nil(err)
	assert.Equal(expectedMovie, foundMovie)
	assert.Equal(expectedRowsAffected, rowsAffected)
	teardown()
}

func TestGetMovieByNonExistingId(t *testing.T) {
	setup()
	store := MovieStore{}
	assert := assert.New(t)
	query := "^SELECT (.+) FROM \"movies\" WHERE \\(id = \\?\\) LIMIT 1$"

	expectedRowsAffected := int64(0)
	expectedMovie := models.Movie{
		Id:     0,
		Name:   "",
		Rating: 0,
	}

	rows := sqlmock.NewRows([]string{})

	mock.ExpectQuery(query).
		WithArgs(999).
		WillReturnRows(rows)

	foundMovie, affectedRows, err := store.GetMovieById(999)

	assert.True(gorm.IsRecordNotFoundError(err))
	assert.Equal(expectedMovie, foundMovie)
	assert.Equal(expectedRowsAffected, affectedRows)
	teardown()
}

func TestCreateMovie(t *testing.T) {
	setup()
	store := MovieStore{}
	assert := assert.New(t)
	insertStatement := "^INSERT INTO \"movies\" \\(\"name\",\"rating\"\\) VALUES \\(\\?,\\?\\)$"

	movieToCreate := models.Movie{
		Name:   "Inception",
		Rating: 5,
	}

	expectedRowsAffected := int64(1)
	expectedMovie := models.Movie{
		Id:     1,
		Name:   "Inception",
		Rating: 5,
	}

	mock.ExpectBegin()
	mock.ExpectExec(insertStatement).
		WithArgs(movieToCreate.Name, movieToCreate.Rating).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	createdMovie, rowsAffected, err := store.CreateMovie(&movieToCreate)

	assert.Nil(err)
	assert.Equal(expectedMovie, createdMovie)
	assert.Equal(expectedRowsAffected, rowsAffected)
	teardown()
}

// TODO: implement
func TestCreateMovieWithInvalidData(t *testing.T) {
	setup()
	teardown()
}

func TestUpdateMovie(t *testing.T) {
	setup()
	store := MovieStore{}
	assert := assert.New(t)
	updateStatement := "^UPDATE \"movies\" " +
		"SET \"id\" = \\?, \"name\" = \\?, \"rating\" = \\? " +
		"WHERE \"movies\"\\.\"id\" = \\? AND \\(\\(id = \\?\\)\\)$"

	movieToUpdate := models.Movie{
		Id:     1,
		Name:   "Inception",
		Rating: 5,
	}

	expectedRowsAffected := int64(1)

	mock.ExpectBegin()
	mock.ExpectExec(updateStatement).
		WithArgs(1, movieToUpdate.Name, movieToUpdate.Rating, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	updatedMovie, rowsAffected, err := store.UpdateMovie(1, &movieToUpdate)

	assert.Nil(err)
	assert.Equal(movieToUpdate, updatedMovie)
	assert.Equal(expectedRowsAffected, rowsAffected)
	teardown()
}

// TODO: implement
func TestUpdateMovieWithInvalidData(t *testing.T) {
	setup()
	teardown()
}

func TestDeleteMovieById(t *testing.T) {
	setup()
	store := MovieStore{}
	assert := assert.New(t)
	deleteStatement := "^DELETE FROM \"movies\" WHERE \"movies\"\\.\"id\" = \\?$"

	expectedRowsAffected := int64(1)

	mock.ExpectBegin()
	mock.ExpectExec(deleteStatement).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	rowsAffected, err := store.DeleteMovieById(1)

	assert.Nil(err)
	assert.Equal(expectedRowsAffected, rowsAffected)
	teardown()
}

func TestDeleteMovieByNonExistingId(t *testing.T) {
	setup()
	store := MovieStore{}
	assert := assert.New(t)
	deleteStatement := "^DELETE FROM \"movies\" WHERE \"movies\"\\.\"id\" = \\?$"

	expectedRowsAffected := int64(0)

	mock.ExpectBegin()
	mock.ExpectExec(deleteStatement).
		WithArgs(999).
		WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectCommit()

	rowsAffected, err := store.DeleteMovieById(999)

	assert.True(gorm.IsRecordNotFoundError(err))
	assert.Equal(expectedRowsAffected, rowsAffected)
	teardown()
}
