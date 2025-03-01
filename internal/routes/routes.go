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
	client.Get("/client", handler.GetClient)
	client.Get("/client/:id", handler.FindClient)
	client.Put("/client/:id", handler.UpdateClient)
	client.Delete("/client/:id", handler.RemoveClient)

	// Role Route
	role := app.Group("/role",middleware.Role("admin")) //Role access
	role.Get("/role", handler.GetRole)
	role.Get("/role/:id", handler.FindRole)
	role.Post("/role/store", handler.CreateRole)
	role.Put("/role/:id", handler.UpdateRole)
	role.Delete("/role/:id", handler.DeleteRole)

	// Permission Route
	permission := app.Group("/permission", middleware.Role("admin")) //Role access
	permission.Get("/permission", handler.GetPermission)
	permission.Get("/permission/:id", handler.FindPermission)
	permission.Post("/permission/store", handler.CreatePermission)
	permission.Put("/permission/:id", handler.UpdatePermission)
	permission.Delete("/permission/:id", handler.DeletePermission)

	// User Route
	app.Get("/account/profile", handler.GetProfile)
	app.Put("/account/update", handler.UpdateAccount)
	app.Delete("/account/delete", handler.DeleteAccount)
	app.Post("/logout", handler.Logout)

	// Category Route
	app.Get("/category", handler.GetCategory)
	app.Get("/category/:id", handler.FindCategory)

	category := app.Group("/category",middleware.Permission("kelola category")) // Permission Access
	category.Post("/store", handler.CreateCategory)
	category.Put("/:id", handler.UpdateCategory)
	category.Delete("/:id", handler.DeleteCategory)

	// Product Route
	app.Get("/product", handler.GetProduct)
	app.Get("/product/:id", handler.FindProduct)

	product := app.Group("/product",middleware.Permission("kelola product")) // Permission Access
	product.Post("/product/store", handler.CreateProduct)
	product.Put("/product/:id", handler.UpdateProduct)
	product.Delete("/product/:id", handler.DeleteProduct)
}
