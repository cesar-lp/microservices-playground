package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "movie-service"
)

var database *gorm.DB

func Connect() *gorm.DB {
	var err error

	fmt.Println("Connecting to database...")
	DB_URL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	database, err = gorm.Open("postgres", DB_URL)

	if err != nil {
		fmt.Printf("Cannot connect to database")
		panic(err)
	}

	fmt.Println("Connected to database")

	return database
}

func Database() *gorm.DB {
	return database
}
