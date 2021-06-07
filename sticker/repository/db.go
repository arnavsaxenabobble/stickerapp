package repository

import (
	"stickerapp/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Create a new connection of the database
func NewDbConn() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DatabaseConfig.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}
