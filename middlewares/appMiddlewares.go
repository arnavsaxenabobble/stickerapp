package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

// Function to register middlewares in echo framework
// This method will register both readymade middlewares and custom middlewares
func RegisterMiddleware(e *echo.Echo) {
	// Middleware for logging
	e.Use(middleware.Logger())
	// Middleware for recovery handling
	// e.Use(middleware.Recover())
	// Middleware for CORS enablement
	if viper.GetBool("ENABLE_CORS") == true {
		e.Use(middleware.CORS())
	}
}
