package models

import (
	"strconv"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/common"
)

// Movie model
type Movie struct {
	ID     int    `gorm:"primary_key;auto_increment" json:"id"`
	Name   string `gorm:"size:255;not null;unique" json:"name"`
	Rating int    `json:"rating"`
}

// Initialize the default values
func (movie *Movie) Initialize() {
	movie.ID = 0
}

// Validate a movie structure
func (movie *Movie) Validate() []FieldError {
	fieldErrors := make([]FieldError, 0, 2)

	if movie.Name == "" {
		fieldErrors = append(fieldErrors, FieldError{
			FieldName:    "name",
			Error:        "Name can't be empty",
			InvalidValue: movie.Name,
		})
	}

	if movie.Rating < 0 || movie.Rating > 5 {
		fieldErrors = append(fieldErrors, FieldError{
			FieldName:    "rating",
			Error:        "Rating must be between 0 - 5",
			InvalidValue: strconv.Itoa(movie.Rating),
		})
	}

	return fieldErrors
}
