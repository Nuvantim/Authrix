package routes

import (
	"api/internal/domain/handlers"
	"api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	//auth Route
	app.Post("/account/register", handler.RegisterAccount)
	app.Post("/login", handler.Login)

	//protected
	app.Use(middleware.Setup())

	// Permission Route
	app.Get("/permission",handler.GetPermission)
	app.Get("/permission/:id",handler.FindPermission)
	app.Post("/permission/store",handler.CreatePermission)
	app.Put("/permission/:id",handler.UpdatePermission)
	app.Delete("/permission/:id",handler.DeletePermission)

	// User Route
	app.Get("/account/profile", handler.GetProfile)
	app.Put("/account/update", handler.UpdateAccount)
	app.Delete("/account/delete", handler.DeleteAccount)
	app.Post("/logout", handler.Logout)

	// Category Route
	app.Get("/category", handler.GetCategory)
	app.Get("/category/:id", handler.FindCategory)
	app.Post("/category/store", handler.CreateCategory)
	app.Put("/category/:id", handler.UpdateCategory)
	app.Delete("/category/:id", handler.DeleteCategory)

	// Product Route
	app.Get("/product", handler.GetProduct)
	app.Get("/product/:id", handler.FindProduct)
	app.Post("/product/store", handler.CreateProduct)
	app.Put("/product/:id", handler.UpdateProduct)
	app.Delete("/product/:id", handler.DeleteProduct)
}
