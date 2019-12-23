package models

import (
	"math/rand"
	"strconv"
)

// Movie model
type Movie struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}

var movies = []Movie{
	Movie{ID: "1", Name: "Inception", Rating: 5},
	Movie{ID: "2", Name: "Interstellar", Rating: 5},
	Movie{ID: "3", Name: "The Dark Knight", Rating: 5},
}

// GetAllMovies returns all movies
func GetAllMovies() []Movie {
	return movies
}

// GetMovieByID returns a Movie for a given ID
func GetMovieByID(id string) Movie {
	pos := 0

	for i, movie := range movies {
		if movie.ID == id {
			pos = i
			break
		}
	}

	return movies[pos]
}

// CreateMovie creates a movie
func CreateMovie(newMovie Movie) Movie {
	newMovie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, newMovie)
	return newMovie
}

// UpdateMovie updates a movie
func UpdateMovie(id string, updatedMovie Movie) Movie {
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

// DeleteMovieByID deletes a movie for a given ID
func DeleteMovieByID(id string) {
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}
}
