package domain

import "gorm.io/gorm"

// Store the actual stickers
type Sticker struct {
	// defaults fields of gorm
	gorm.Model
	// the sticker Name
	StickerName string
	// the keyword used to get sticker
	Trending int
	// the actual sticker
	Clicks int
}

// CDNs of stickers
type CDN struct {
	// CDN to be returned
	Images []string `json:"images"`
}
