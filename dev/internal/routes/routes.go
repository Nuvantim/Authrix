package routes

import (
	"api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", handler.Home)
	app.Post("/auth/send-otp", handler.SendOTP)
	app.Post("/auth/register", handler.Register)
	app.Post("auth/login", handler.Login)
	app.Post("auth/reset-password", handler.ResetPassword)

	// user
	app.Get("/account/profile", handler.GetProfile)
	app.Put("/account/update", handler.UpdateProfile)
}
