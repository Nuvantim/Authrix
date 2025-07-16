package routes

import (
	"api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", handler.Home)
	app.Post("/send/otp", handler.SendOTP)
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	app.Post("/update/password", handler.UpdatePassword)
}
