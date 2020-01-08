package database

import (
	"log"

	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
)

var movies = []models.Movie{
	models.Movie{Name: "Inception", Rating: 5},
	models.Movie{Name: "Interstellar", Rating: 5},
	models.Movie{Name: "The Dark Knight", Rating: 5},
}

// Load seeds to a given database
func Load() {
	err := database.Debug().DropTableIfExists(&models.Movie{}).Error

	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}

	err = database.Debug().AutoMigrate(&models.Movie{}).Error

	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	for i, _ := range movies {
		err = database.Debug().Model(&models.Movie{}).Create(&movies[i]).Error

		if err != nil {
			log.Fatalf("Cannot seed movies table: %v", err)
		}
	}
}
