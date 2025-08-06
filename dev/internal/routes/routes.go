package routes

import (
	"api/internal/app/handlers"
	"api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", handler.Home)

	app.Use(middleware.Setup())

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

	// client
	client := app.Group("/client")
	client.Get("/", handler.ListClient)
	client.Get("/:id", handler.GetClient)
	client.Put("/update/:id", handler.UpdateClient)
	client.Delete("/delete/:id", handler.DeleteClient)

	// permission
	permission := app.Group("/permission")
	permission.Get("/", handler.ListPermission)
	permission.Get("/:id", handler.GetPermission)
	permission.Post("/store", handler.CreatePermission)
	permission.Put("/update/:id", handler.UpdatePermission)
	permission.Delete("/delete/:id", handler.DeletePermission)

	// role
	role := app.Group("/role")
	role.Get("/", handler.ListRole)
	role.Get("/:id", handler.GetRole)
	role.Post("/store", handler.CreateRole)
	role.Put("/update/:id", handler.UpdateRole)
	role.Delete("/delete/:id", handler.DeleteRole)

}
