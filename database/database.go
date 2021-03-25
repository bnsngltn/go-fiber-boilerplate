package database

import (
	"fmt"
	"log"

	"github.com/bnsngltn/go-fiber-boilerplate/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DB is the database instance
var DB *gorm.DB

//Connect attempts to our database
func Connect() error {
	log.Println("Attempting to connect to the database...")

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	return nil
}
