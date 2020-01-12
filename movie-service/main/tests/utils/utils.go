package utils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

// DBConnectionMock structure.
type DBConnectionMock struct {
	db   *sql.DB
	Gorm *gorm.DB
	Mock sqlmock.Sqlmock
}

// WriteBody writes content to a HTTP request body.
func WriteBody(content interface{}) io.Reader {
	b, _ := json.Marshal(content)
	return ioutil.NopCloser(bytes.NewBuffer(b))
}

// MockDBConnection mocks a db connection and returns it.
func MockDBConnection() DBConnectionMock {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatalf("Error '%s' when opening a stub database connection", err)
	}

	gorm, err := gorm.Open("sqlmock", db)

	if err != nil {
		log.Fatal("Could not mock a gorm connection")
	}

	return DBConnectionMock{db: db, Gorm: gorm, Mock: mock}
}

// AssertAllExpectationsWereMet asserts all expectations were met on a mocked db connection.
func (dbConnection DBConnectionMock) AssertAllExpectationsWereMet() {
	if err := dbConnection.Mock.ExpectationsWereMet(); err != nil {
		log.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

// Close the mocked db connection.
func (dbConnection *DBConnectionMock) Close() {
	dbConnection.db.Close()
	dbConnection.Gorm.Close()
}
