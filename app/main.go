package main

import (
	"log"
	"os"

	h "stickerapp/sticker/delivery/http"
	m "stickerapp/sticker/delivery/http/middlewares"

	"github.com/labstack/echo/v4"
)

// Global Variables
var (
	e *echo.Echo
)

// Init function to initialize an instance of a Echo Server and configure middlewares
func init() {
	// Create and configure the instance of Echo with middlewares for creating the web-app
	e = echo.New()
	// Register middlewares to the echo server
	m.RegisterMiddleware(e)
	// Register the routes/end-points and their corresponding function handlers
	h.RegisterHandlers(e)
}

// Main Function: Start point of the flow
func main() {
	// Start the web application on a random available port
	port := os.Getenv("PORT")
	log.Fatal(e.Start(":" + port))
}
