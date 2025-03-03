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
	app.Get("/test", handler.Test)

	// Client Route
	client := app.Group("/client", middleware.Role("admin")) //Role access
	client.Get("", handler.GetClient)
	client.Get("/:id", handler.FindClient)
	client.Put("/:id", handler.UpdateClient)
	client.Delete("/:id", handler.RemoveClient)

	// Role Route
	role := app.Group("/role", middleware.Role("admin")) //Role access
	role.Get("", handler.GetRole)
	role.Get("/:id", handler.FindRole)
	role.Post("/store", handler.CreateRole)
	role.Put("/:id", handler.UpdateRole)
	role.Delete("/:id", handler.DeleteRole)

	// Permission Route
	permission := app.Group("/permission", middleware.Role("admin")) //Role access
	permission.Get("", handler.GetPermission)
	permission.Get("/:id", handler.FindPermission)
	permission.Post("/store", handler.CreatePermission)
	permission.Put("/:id", handler.UpdatePermission)
	permission.Delete("/:id", handler.DeletePermission)

	// User Route
	app.Get("/account/profile", handler.GetProfile)
	app.Put("/account/update", handler.UpdateAccount)
	app.Delete("/account/delete", handler.DeleteAccount)
	app.Post("/logout", handler.Logout)

	// Category Route
	app.Get("/category", handler.GetCategory)
	app.Get("/category/:id", handler.FindCategory)

	category := app.Group("/category", middleware.Permission("kelola category")) // Permission Access
	category.Post("/store", handler.CreateCategory)
	category.Put("/:id", handler.UpdateCategory)
	category.Delete("/:id", handler.DeleteCategory)

	// Product Route
	app.Get("/product", handler.GetProduct)
	app.Get("/product/:id", handler.FindProduct)

	product := app.Group("/product", middleware.Permission("kelola product")) // Permission Access
	product.Post("/product/store", handler.CreateProduct)
	product.Put("/product/:id", handler.UpdateProduct)
	product.Delete("/product/:id", handler.DeleteProduct)
}
