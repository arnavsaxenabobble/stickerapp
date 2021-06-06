package http

import (
	"net/http"

	r "stickerapp/sticker/repository"

	"github.com/labstack/echo/v4"
)

// Function handler for GET - http://<HOST>:<PORT>/ to check application liveliness
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "The server is live...")
}

// Data field of response
type response struct {
	Images []string `json:"images"`
}

// response variable
var res response

// Function handler for GET - http://<HOST>:<PORT>/v1/trendingStickers to get trending sticker(s)
func getTrendingStickers(c echo.Context) error {
	res.Images = nil
	// Fetch the sticker id if present
	name := c.QueryParam("name")
	// Fetch the pagination limit if present
	pageLimit := c.QueryParam("limit")
	// // Fetch the stickers info from database and return the CDN of the sticker(s)
	stickerInfo := r.FetchStickerInfo(name, pageLimit)
	// Fetch CDN endpoint
	cdnUrl := r.CdnUrl()
	// Calculate complete CDN path of Stickers
	stickers := stickerInfo[:0]
	for _, sticker := range stickerInfo {
		res.Images = append(res.Images, cdnUrl+"/"+sticker.StickerName)
	}

	if len(stickers) < 0 {
		return c.JSON(http.StatusNotFound, "Sticker not found")
	}

	return c.JSON(http.StatusOK, res)
}
