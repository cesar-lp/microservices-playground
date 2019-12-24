package handlers

import (
	"math/rand"
	"strconv"

	. "github.com/cesar-lp/microservices-playground/movie-service/main/models"
)

var movies = []Movie{
	Movie{ID: "1", Name: "Inception", Rating: 5},
	Movie{ID: "2", Name: "Interstellar", Rating: 5},
	Movie{ID: "3", Name: "The Dark Knight", Rating: 5},
}

// GetAll returns all movies
func GetAll() []Movie {
	return movies
}

// Get returns a Movie for a given ID
func Get(id string) Movie {
	pos := 0

	for i, movie := range movies {
		if movie.ID == id {
			pos = i
			break
		}
	}

	return movies[pos]
}

// Save a movie
func Save(newMovie Movie) Movie {
	newMovie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, newMovie)
	return newMovie
}

// Update a movie
func Update(id string, updatedMovie Movie) Movie {
	pos := 0

	for i, movie := range movies {
		if movie.ID == id {
			pos = i
			movies[i] = updatedMovie
			break
		}
	}

	return movies[pos]
}

// Delete a movie for a given ID
func Delete(id string) {
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}
}
