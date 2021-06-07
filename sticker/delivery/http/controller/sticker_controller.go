package controller

import (
	"encoding/json"
	"net/http"
	"stickerapp/sticker/delivery/http/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

// interface for sticker app controller
type StickerController interface {
	FindAll(ctx echo.Context) error
	FindByName(ctx echo.Context) error
}

// controller struct
type controller struct {
	service service.StickerService
}

// constructor for controller
func New(service service.StickerService) StickerController {
	return &controller{
		service: service,
	}
}

// find all method to use service and repository to find all stickers
func (c *controller) FindAll(ctx echo.Context) error {
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		panic(err)
	}
	result := c.service.FindAll(limit)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	ctx.Response().WriteHeader(http.StatusOK)
	json.NewEncoder(ctx.Response()).Encode(result)
	return nil
}

// find by name method to use service and repository to find sticker by name
func (c *controller) FindByName(ctx echo.Context) error {
	name := ctx.Param("name")
	if len(ctx.Param("limit")) > 0 {
		limit, err := strconv.Atoi(ctx.QueryParam("limit"))
		if err != nil {
			panic(err)
		}
		result := c.service.FindByName(name, limit)
		ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		ctx.Response().WriteHeader(http.StatusOK)
		json.NewEncoder(ctx.Response()).Encode(result)
	} else {
		result := c.service.FindByName(name, 0)
		ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		ctx.Response().WriteHeader(http.StatusOK)
		json.NewEncoder(ctx.Response()).Encode(result)
	}
	return nil
}
