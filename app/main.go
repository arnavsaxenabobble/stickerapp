package main

import (
	"log"
	"os"

	c "stickerapp/config"
	h "stickerapp/sticker/delivery/http"
	m "stickerapp/sticker/delivery/http/middlewares"
	r "stickerapp/sticker/repository"
	st "stickerapp/sticker/storage"

	"github.com/labstack/echo/v4"
)

// Global Variables
var (
	e *echo.Echo
	d *echo.Group
)

// Init function to initialize an instance of a Echo Server and configure middlewares
func init() {
	// Create and configure the instance of Echo with middlewares for creating the web-app
	e = echo.New()
	// Read application configuration from configuration file
	err := c.GetAppConfig()
	if err != nil {
		e.Logger.Fatal("Failed to read configuration %s", err)
	}
	// Register middlewares to the echo server
	m.RegisterMiddleware(e)
	// Intialise echo group for documentation routes
	d = e.Group("")
	// Register the routes/end-points and their corresponding function handlers
	h.RegisterHandlers(e)
}

// Main Function: Start point of the flow
func main() {
	// Start the web application on a random available port
	port := os.Getenv("PORT")
	log.Fatal(e.Start(":" + port))
	// Get database configuration and perform initial migration
	err := c.GetDatabaseConfig()
	if err != nil {
		e.Logger.Fatal("Failed to load database configuration %s", err)
	}
	initialMigration()
}

// InitialMigration for sticker with db.AutoMigrate
func initialMigration() {
	db := st.NewDbConn()
	// Migrate the schema
	sticker := r.Sticker{}
	e.Logger.Debug("Migrating...")
	db.AutoMigrate(&sticker)
	e.Logger.Debug("Migrated...")
}
