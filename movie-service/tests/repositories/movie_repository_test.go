package repositories_tests

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/cesar-lp/microservices-playground/movie-service/main/repositories"
	"github.com/cesar-lp/microservices-playground/movie-service/tests/utils"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func setupTest(dbConn utils.DBConnectionMock, t *testing.T) (repositories.MovieRepository, *assert.Assertions) {
	return repositories.MovieRepositoryImpl(dbConn.Gorm), assert.New(t)
}

func TestGetAllMovies_ReturnsEmpty(t *testing.T) {
	dbConnection := utils.MockDBConnection()
	defer dbConnection.Close()
	repository, assert := setupTest(dbConnection, t)
	query := "^SELECT (.+) FROM \"movies\"$"

	expectedRowsAffected := int64(0)

	dbConnection.Mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{}))

	foundMovies, rowsAffected, err := repository.GetAllMovies()

	assert.Nil(err)
	assert.Len(foundMovies, 0)
	assert.Equal(expectedRowsAffected, rowsAffected)
	dbConnection.AssertAllExpectationsWereMet()
}

func TestGetAllMovies(t *testing.T) {
	dbConnection := utils.MockDBConnection()
	defer dbConnection.Close()
	repository, assert := setupTest(dbConnection, t)
	query := "^SELECT (.+) FROM \"movies\"$"

	expectedRowsAffected := int64(3)
	expectedMovies := []models.Movie{
		models.Movie{Name: "Inception", Rating: 5},
		models.Movie{Name: "Interstellar", Rating: 5},
		models.Movie{Name: "The Dark Knight", Rating: 5},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "rating"})

	for _, m := range expectedMovies {
		rows.AddRow(m.Id, m.Name, m.Rating)
	}

	dbConnection.Mock.ExpectQuery(query).WillReturnRows(rows)

	foundMovies, rowsAffected, err := repository.GetAllMovies()

	assert.Nil(err)
	assert.Len(foundMovies, len(expectedMovies))
	assert.Equal(expectedMovies, foundMovies)
	assert.Equal(expectedRowsAffected, rowsAffected)
	dbConnection.AssertAllExpectationsWereMet()
}

func TestGetMovieById(t *testing.T) {
	dbConnection := utils.MockDBConnection()
	defer dbConnection.Close()
	repository, assert := setupTest(dbConnection, t)
	query := "^SELECT (.+) FROM \"movies\" WHERE \\(id = \\?\\) LIMIT 1$"

	expectedRowsAffected := int64(1)
	expectedMovie := models.Movie{
		Id:     1,
		Name:   "Inception",
		Rating: 5,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "rating"}).
		AddRow(expectedMovie.Id, expectedMovie.Name, expectedMovie.Rating)

	dbConnection.Mock.ExpectQuery(query).
		WithArgs(1).
		WillReturnRows(rows)

	foundMovie, rowsAffected, err := repository.GetMovieById(1)

	assert.Nil(err)
	assert.Equal(expectedMovie, foundMovie)
	assert.Equal(expectedRowsAffected, rowsAffected)
	dbConnection.AssertAllExpectationsWereMet()
}

func TestGetMovieById_ReturnsRecordNotFoundError(t *testing.T) {
	dbConnection := utils.MockDBConnection()
	defer dbConnection.Close()
	repository, assert := setupTest(dbConnection, t)
	query := "^SELECT (.+) FROM \"movies\" WHERE \\(id = \\?\\) LIMIT 1$"

	expectedRowsAffected := int64(0)
	expectedMovie := models.Movie{
		Id:     0,
		Name:   "",
		Rating: 0,
	}

	rows := sqlmock.NewRows([]string{})

	dbConnection.Mock.ExpectQuery(query).
		WithArgs(999).
		WillReturnRows(rows)

	foundMovie, affectedRows, err := repository.GetMovieById(999)

	assert.True(gorm.IsRecordNotFoundError(err))
	assert.Equal(expectedMovie, foundMovie)
	assert.Equal(expectedRowsAffected, affectedRows)
	dbConnection.AssertAllExpectationsWereMet()
}

func TestCreateMovie(t *testing.T) {
	dbConnection := utils.MockDBConnection()
	defer dbConnection.Close()
	repository, assert := setupTest(dbConnection, t)
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

	dbConnection.Mock.ExpectBegin()
	dbConnection.Mock.ExpectExec(insertStatement).
		WithArgs(movieToCreate.Name, movieToCreate.Rating).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbConnection.Mock.ExpectCommit()

	createdMovie, rowsAffected, err := repository.CreateMovie(&movieToCreate)

	assert.Nil(err)
	assert.Equal(expectedMovie, createdMovie)
	assert.Equal(expectedRowsAffected, rowsAffected)
	dbConnection.AssertAllExpectationsWereMet()
}

func TestCreateMovie_ReturnsDuplicatedKeyError(t *testing.T) {
	dbConnection := utils.MockDBConnection()
	defer dbConnection.Close()
	repository, assert := setupTest(dbConnection, t)
	insertStatement := "^INSERT INTO \"movies\" \\(\"name\",\"rating\"\\) VALUES \\(\\?,\\?\\)$"
	duplicatedKeyError := "pq: duplicate key value violates unique constraint \"movies_name_key\""

	movieToCreate := models.Movie{
		Name:   "Duplicated movie name",
		Rating: 5,
	}

	expectedRowsAffected := int64(1)
	expectedMovie := models.Movie{
		Id:     1,
		Name:   "Duplicated movie name",
		Rating: 5,
	}

	dbConnection.Mock.ExpectBegin()
	dbConnection.Mock.ExpectExec(insertStatement).
		WithArgs(movieToCreate.Name, movieToCreate.Rating).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbConnection.Mock.ExpectCommit()

	createdMovie, rowsAffected, err := repository.CreateMovie(&movieToCreate)

	assert.Nil(err)
	assert.Equal(expectedRowsAffected, rowsAffected)
	assert.Equal(expectedMovie, createdMovie)

	expectedRowsAffected = int64(0)
	duplicatedMovieToCreate := models.Movie{
		Name:   "Duplicated movie name",
		Rating: 5,
	}

	dbConnection.Mock.ExpectBegin()
	dbConnection.Mock.ExpectExec(insertStatement).
		WithArgs(duplicatedMovieToCreate.Name, duplicatedMovieToCreate.Rating).
		WillReturnResult(sqlmock.NewResult(1, 0)).
		WillReturnError(errors.New(duplicatedKeyError))
	dbConnection.Mock.ExpectRollback()

	_, rowsAffected, err = repository.CreateMovie(&duplicatedMovieToCreate)

	assert.Equal(duplicatedKeyError, err.Error())
	assert.Equal(expectedRowsAffected, rowsAffected)
	dbConnection.AssertAllExpectationsWereMet()
}

func TestUpdateMovie(t *testing.T) {
	dbConnection := utils.MockDBConnection()
	defer dbConnection.Close()
	repository, assert := setupTest(dbConnection, t)
	updateStatement := "^UPDATE \"movies\" " +
		"SET \"id\" = \\?, \"name\" = \\?, \"rating\" = \\? " +
		"WHERE \"movies\"\\.\"id\" = \\? AND \\(\\(id = \\?\\)\\)$"

	movieToUpdate := models.Movie{
		Id:     1,
		Name:   "Inception",
		Rating: 5,
	}

	expectedRowsAffected := int64(1)

	dbConnection.Mock.ExpectBegin()
	dbConnection.Mock.ExpectExec(updateStatement).
		WithArgs(1, movieToUpdate.Name, movieToUpdate.Rating, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbConnection.Mock.ExpectCommit()

	updatedMovie, rowsAffected, err := repository.UpdateMovie(1, &movieToUpdate)

	assert.Nil(err)
	assert.Equal(movieToUpdate, updatedMovie)
	assert.Equal(expectedRowsAffected, rowsAffected)
	dbConnection.AssertAllExpectationsWereMet()
}

func TestDeleteMovieById(t *testing.T) {
	dbConnection := utils.MockDBConnection()
	defer dbConnection.Close()
	repository, assert := setupTest(dbConnection, t)
	deleteStatement := "^DELETE FROM \"movies\" WHERE \"movies\"\\.\"id\" = \\?$"

	expectedRowsAffected := int64(1)

	dbConnection.Mock.ExpectBegin()
	dbConnection.Mock.ExpectExec(deleteStatement).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbConnection.Mock.ExpectCommit()

	rowsAffected, err := repository.DeleteMovieById(1)

	assert.Nil(err)
	assert.Equal(expectedRowsAffected, rowsAffected)
	dbConnection.AssertAllExpectationsWereMet()
}

func TestDeleteMovieById_ReturnsRecordNotFound(t *testing.T) {
	dbConnection := utils.MockDBConnection()
	defer dbConnection.Close()
	repository, assert := setupTest(dbConnection, t)
	deleteStatement := "^DELETE FROM \"movies\" WHERE \"movies\"\\.\"id\" = \\?$"

	expectedRowsAffected := int64(0)

	dbConnection.Mock.ExpectBegin()
	dbConnection.Mock.ExpectExec(deleteStatement).
		WithArgs(999).
		WillReturnResult(sqlmock.NewResult(1, 0))
	dbConnection.Mock.ExpectCommit()

	rowsAffected, err := repository.DeleteMovieById(999)

	assert.True(gorm.IsRecordNotFoundError(err))
	assert.Equal(expectedRowsAffected, rowsAffected)
	dbConnection.AssertAllExpectationsWereMet()
}
