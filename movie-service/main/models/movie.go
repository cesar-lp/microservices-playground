package models

import (
	"strconv"

	"github.com/cesar-lp/microservices-playground/movie-service/main/common"
)

// Movie structure.
type Movie struct {
	Id     int    `gorm:"primary_key;auto_increment" json:"id"`
	Name   string `gorm:"size:255;not null;unique" json:"name"`
	Rating int    `json:"rating"`
}

// Initialize the default values of a given movie.
func (movie *Movie) Initialize() {
	movie.Id = 0
}

// Validate a movie structure and returns all the errors encountered.
func (movie *Movie) Validate() []common.FieldError {
	fieldErrors := make([]common.FieldError, 0, 2)

	if movie.Name == "" {
		fieldErrors = append(fieldErrors,
			common.NewFieldError("name", "Name is required", movie.Name))
	}

	if movie.Rating < 0 || movie.Rating > 5 {
		fieldErrors = append(fieldErrors,
			common.NewFieldError("rating", "Rating must be betweent 0 - 5", strconv.Itoa(movie.Rating)))
	}

	return fieldErrors
}
