package repository

import (
	"stickerapp/config"
	"stickerapp/sticker/domain"

	"gorm.io/gorm"
)

// interface of sticker repository
type StickerRepository interface {
	FindByName(name string, limit int) []domain.Sticker
	FindAll(limit int) []domain.Sticker
}

// struct for repository
type database struct {
	connection *gorm.DB
}

// creates a new connection with sticker repository
func NewStickerRepository() StickerRepository {
	config.GetDatabaseConfig()
	db := NewDbConn()
	return &database{
		connection: db,
	}
}

// Find a particular sticker by its name
func (db *database) FindByName(name string, limit int) []domain.Sticker {
	var stickers []domain.Sticker
	db.connection.Where("sticker_name = ?", name).Find(&stickers)
	return stickers
}

// Searches all the stickers and return results sorted by trending bound by a limit
func (db *database) FindAll(limit int) []domain.Sticker {
	var stickers []domain.Sticker
	db.connection.Limit(limit).Order("trending").Find(&stickers)
	return stickers
}
