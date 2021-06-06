package storage

import (
	c "stickerapp/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database Connection variable
var DB *gorm.DB

func NewDbConn() *gorm.DB {
	var err error
	DB, err = gorm.Open(postgres.Open(c.DatabaseConfig.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
