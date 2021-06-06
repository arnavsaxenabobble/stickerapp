package http

import "github.com/labstack/echo/v4"

// Function handler to register routes of the application
func RegisterHandlers(e *echo.Echo) {
	e.GET("/v1/health", healthCheck)
	e.GET("/v1/trendingStickers", getTrendingStickers)
}
