package handlers

import (
	"errors"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/common"
	. "github.com/cesar-lp/microservices-playground/movie-service/main/database"
	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	. "github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/jinzhu/gorm"
)

// GetAll returns all movies
func GetAll() HandlerResponse {
	movies := []models.Movie{}

	var err = Database().Debug().Model(&models.Movie{}).Limit(50).Find(&movies).Error

	if err != nil {
		return InternalServerError(err)
	}
	return Ok(movies)
}

// Get returns a Movie for a given ID
func Get(id string) HandlerResponse {
	movie := models.Movie{}

	var err = Database().Debug().Model(&models.Movie{}).Where("id = ?", id).Take(&movie).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + id))
		}
		return InternalServerError(err)
	}
	return Ok(movie)
}

// Save a movie
func Save(newMovie Movie) HandlerResponse {
	fieldErrors := newMovie.Validate()

	if len(fieldErrors) > 0 {
		return UnprocessableEntity(fieldErrors)
	}

	newMovie.Initialize()

	var err = Database().Debug().Model(&models.Movie{}).Create(&newMovie).Error

	if err != nil {
		return InternalServerError(err)
	}
	return Created(newMovie)
}

// Update a movie
func Update(id string, updatedMovie Movie) HandlerResponse {
	fieldErrors := updatedMovie.Validate()

	if len(fieldErrors) > 0 {
		return UnprocessableEntity(fieldErrors)
	}

	var err = Database().Debug().Model(&models.Movie{}).Where("id = ?", id).Update(&updatedMovie).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + id))
		}
		return InternalServerError(err)
	}
	return Ok(&updatedMovie)
}

// Delete a movie for a given ID
func Delete(id string) HandlerResponse {
	var err = Database().Debug().Model(&models.Movie{}).Where("id = ?", id).Take(&Movie{}).Delete(&Movie{}).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFound(errors.New("Movie not found for id " + id))
		}
		return InternalServerError(err)
	}
	return NoContent()
}
