package routes

import (
	"api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/hello", handler.Hello)
}
