package server

import (
	"fmt"

	"github.com/cesar-lp/microservices-playground/movie-service/main/app/config"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/models"
	"github.com/cesar-lp/microservices-playground/movie-service/main/app/server/seeds"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Database structure.
type Database struct {
	instance       *gorm.DB
	host           string
	port           int
	user           string
	password       string
	name           string
	loggingEnabled bool
}

func setupDB(db config.DBConfig) Database {
	return Database{
		instance:       &gorm.DB{},
		host:           db.Host,
		port:           db.Port,
		user:           db.User,
		password:       db.Password,
		name:           db.Name,
		loggingEnabled: db.LoggingEnabled,
	}
}

// Connect establishes a connection to a database using the provided values.
func (db *Database) Connect() {
	var err error

	log.Infof("Connecting to database %s...", db.name)
	DB_URL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.host, db.port, db.user, db.password, db.name)

	database, err := gorm.Open("postgres", DB_URL)

	if err != nil {
		log.Panicf("Cannot connect to database: %s", err)
	}

	if db.loggingEnabled {
		database = database.Debug()
	}

	db.instance = database
	log.Info("Connected to database")
}

// LoadSeeds migrates and loads into database existing models.
func (db *Database) LoadSeeds() {
	log.Info("Loading seeds...")
	err := db.instance.Debug().DropTableIfExists(&models.Movie{}).Error

	if err != nil {
		log.Panicf("Cannot drop table: %v", err)
	}

	err = db.instance.Debug().AutoMigrate(&models.Movie{}).Error

	if err != nil {
		log.Panicf("Cannot migrate table: %v", err)
	}

	for i, _ := range seeds.Movies {
		err = db.instance.Debug().Model(&models.Movie{}).Create(&seeds.Movies[i]).Error

		if err != nil {
			log.Panicf("Cannot seed movies table: %v", err)
		}
	}
	log.Info("Seeds loaded")
}
