package database

import (
	"fmt"

	"github.com/blauwiggle/go-rest-api/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and creates the table
func MigrateDB(db *gorm.DB) error {
	fmt.Println("Migrating database")

	if err := db.AutoMigrate(&comment.Comment{}).Error; err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	return nil
}
