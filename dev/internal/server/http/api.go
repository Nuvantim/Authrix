package http

import (
	"api/config"
	"api/database"
	"api/internal/routes"
	"github.com/gofiber/fiber/v2"
)

// ServerGo initializes and returns a Fiber app instance
func ServerGo() *fiber.App {
	// Start Fiber APP
	app := fiber.New(config.FiberConfig())

	// Security Configuration
	config.SecurityConfig(app)

	// Set up all routes
	routes.Setup(app)

	// Start Database Connection
	database.InitDB()

	return app
}
