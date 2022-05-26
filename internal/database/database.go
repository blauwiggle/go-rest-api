package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDatabase - creates a new database connection
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbName, dbPassword)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Failed to connect to database")
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		fmt.Println("Failed to ping database")
		return nil, err
	}

	return db, nil
}
