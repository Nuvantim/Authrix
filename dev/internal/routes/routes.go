package routes

import (
	"api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", handler.Home)
	auth := app.Group("/auth")
	auth.Post("/send-otp", handler.SendOTP)
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Post("/reset-password", handler.ResetPassword)

	// user
	account := app.Group("/account")
	account.Get("/profile", handler.GetProfile)
	account.Put("/update", handler.UpdateAccount)
	account.Delete("/delete", handler.DeleteAccount)
}
