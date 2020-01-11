package server

import (
	"fmt"
	"log"

	"github.com/cesar-lp/microservices-playground/movie-service/main/models"
	"github.com/cesar-lp/microservices-playground/movie-service/main/server/seeds"
	"github.com/jinzhu/gorm"
)

// Database structure.
type Database struct {
	instance *gorm.DB
	host     string
	port     int
	user     string
	password string
	name     string
	log      bool
}

func setupDB(host string, port int, user, password, name string, log bool) Database {
	return Database{
		instance: &gorm.DB{},
		host:     host,
		port:     port,
		user:     user,
		password: password,
		name:     name,
		log:      log,
	}
}

// Connect establishes a connection to a database using the provided values.
func (db *Database) Connect() {
	var err error

	fmt.Println("Connecting to database...")
	DB_URL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.host, db.port, db.user, db.password, db.name)

	database, err := gorm.Open("postgres", DB_URL)

	if err != nil {
		fmt.Printf("Cannot connect to database")
		panic(err)
	}

	if db.log {
		database = database.Debug()
	}

	db.instance = database
	fmt.Println("Connected to database")
}

// LoadSeeds migrates and loads into database existing models.
func (db *Database) LoadSeeds() {
	err := db.instance.Debug().DropTableIfExists(&models.Movie{}).Error

	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}

	err = db.instance.Debug().AutoMigrate(&models.Movie{}).Error

	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	for i, _ := range seeds.Movies {
		err = db.instance.Debug().Model(&models.Movie{}).Create(&seeds.Movies[i]).Error

		if err != nil {
			log.Fatalf("Cannot seed movies table: %v", err)
		}
	}
}
