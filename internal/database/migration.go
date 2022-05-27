package database

import (
	"github.com/blauwiggle/go-rest-api/internal/comment"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// MigrateDB - migrates our database and creates the table
func MigrateDB(db *gorm.DB) error {
	log.Info("Migrating database")

	if err := db.AutoMigrate(&comment.Comment{}).Error; err != nil {
		log.Error("Failed to migrate database")
		return err
	}

	return nil
}
