package http

import (
	"encoding/json"
	"net/http"

	r "stickerapp/repository"

	"github.com/labstack/echo/v4"
)

// Function handler for GET - http://<HOST>:<PORT>/ to check application liveliness
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "The server is live...")
}

// Data field of response
type Sticker struct {
	Stickers []string `json:"stickers"`
}

// Function handler for GET - http://<HOST>:<PORT>/v1/trendingStickers to get trending sticker(s)
func getTrendingStickers(c echo.Context) error {
	// Fetch the sticker id if present
	stickerName := c.QueryParam("name")
	// Fetch the pagination limit if present
	pageLimit := c.QueryParam("limit")

	// // Fetch the stickers info from database and return the CDN of the sticker(s)
	stickerNames, err := r.FetchStickerInfo(stickerName, pageLimit)
	if err != nil {
		panic(err)
	}

	// // Fetch CDN endpoint
	// cdnUrl := r.CdnUrl()
	// // Calculate complete CDN path of Stickers
	// stickers := stickerNames[:0]
	// for _, sticker := range stickerNames {
	// 	stickers = append(stickers, cdnUrl+"/"+sticker)
	// }

	// if len(stickers) < 0 {
	// 	return c.JSON(http.StatusNotFound, "Sticker not found")
	// }
	// response := &Sticker{
	// 	Stickers: stickers,
	// }

	return c.JSON(http.StatusOK, stickerNames)
}

// Function handler for PUT - http://<HOST>:<PORT>/v1/trendingStickers to update trending sticker(s)
func updateTrendingStickers(c echo.Context) error {
	response, err := json.Marshal("Update Sticker...")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)
}

// Function handler for POST - http://<HOST>:<PORT>/v1/sticker to add sticker(s)
func addSticker(c echo.Context) error {
	response, err := json.Marshal("Adding Sticker...")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)
}

// Function handler for DELETE - http://<HOST>:<PORT>/v1/sticker to delete sticker(s)
func deleteSticker(c echo.Context) error {
	response, err := json.Marshal("Update Sticker...")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)
}
