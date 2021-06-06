package repository

import (
	"strconv"

	"gorm.io/driver/postgres"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

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

var stickers []Sticker

// Fetches the info of sticker from DB based on stickerName and limit
func FetchStickerInfo(stickerName string, lim string) ([]Sticker, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=gorm password=gorm dbname=sticker port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Check limit
	var limit int
	if len(lim) != 0 {
		limit, err = strconv.Atoi(lim)
		if err != nil {
			panic(err)
		}
	} else {
		limit = 1
	}
	if len(stickerName) != 0 {
		results := db.Where("sticker_name = ?", stickerName).Limit(limit).Find(&stickers)
		if results.Error != nil {
			panic(results.Error)
		}
		return stickers, nil
	} else {
		results := db.Limit(limit).Find(&stickers)
		if results.Error != nil {
			panic(results.Error)
		}
		return stickers, nil
	}
}

// Fetches the URL of Image CDN
func CdnUrl() string {
	u := viper.GetString("CDN_URL")
	if len(u) == 0 {
		panic("Unable to find CDN URL")
	}
	s := viper.GetString("CDN_ID")
	if len(s) == 0 {
		panic("Unable to find CDN ID")
	}
	// return endpoint of CDN
	return u + s
}
