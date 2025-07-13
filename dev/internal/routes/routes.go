package routes

import (
	"api/internal/domain/handlers"
	// "api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/hello", handler.Hello)
}