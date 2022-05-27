package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

// NewDatabase - creates a new database connection
func NewDatabase() (*gorm.DB, error) {
	log.Info("Setting up new database connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUsername, dbName, dbPassword, sslMode)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Error("Failed to connect to database")
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		log.Error("Failed to ping database")
		return nil, err
	}

	return db, nil
}
