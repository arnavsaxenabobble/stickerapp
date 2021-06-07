package http

import (
	"stickerapp/sticker/delivery/http/controller"
	"stickerapp/sticker/delivery/http/service"
	"stickerapp/sticker/repository"

	"github.com/labstack/echo/v4"
)

// dependency injection to ensure segregation of layers
var (
	stickerRepository repository.StickerRepository = repository.NewStickerRepository()
	stickerService    service.StickerService       = service.New(stickerRepository)
	stickerController controller.StickerController = controller.New(stickerService)
)

// Function handler to register routes of the application
func RegisterHandlers(e *echo.Echo) {
	// endpoint to find all stickers
	e.GET("/v1/trendingStickers", stickerController.FindAll)
	// endpoint to find sticker based on sticker name
	e.GET("/v1/trendingStickers/:name", stickerController.FindByName)
}
