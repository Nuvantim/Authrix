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
	/*
		Role Admin
	*/
	app.Get("/client", middleware.Role("admin"), handler.GetClient)
	app.Get("/client/:id",middleware.Role("admin"), handler.FindClient)
	app.Put("/client/:id",middleware.Role("admin"), handler.UpdateClient)
	app.Delete("/client/:id",middleware.Role("admin"), handler.RemoveClient)

	// Role Route
	app.Get("/role", middleware.Role("admin"), handler.GetRole)
	app.Get("/role/:id", middleware.Role("admin"), handler.FindRole)
	app.Post("/role/store", middleware.Role("admin"), handler.CreateRole)
	app.Put("/role/:id", middleware.Role("admin"), handler.UpdateRole)
	app.Delete("/role/:id", middleware.Role("admin"), handler.DeleteRole)

	// Permission Route
	app.Get("/permission",middleware.Role("admin"), handler.GetPermission)
	app.Get("/permission/:id",middleware.Role("admin"), handler.FindPermission)
	app.Post("/permission/store",middleware.Role("admin"), handler.CreatePermission)
	app.Put("/permission/:id",middleware.Role("admin"), handler.UpdatePermission)
	app.Delete("/permission/:id",middleware.Role("admin"), handler.DeletePermission)

	// User Route
	app.Get("/account/profile", handler.GetProfile)
	app.Put("/account/update", handler.UpdateAccount)
	app.Delete("/account/delete", handler.DeleteAccount)
	app.Post("/logout", handler.Logout)

	// Category Route
	app.Get("/category", handler.GetCategory)
	app.Get("/category/:id", handler.FindCategory)
	app.Post("/category/store",middleware.Permission("kelola category"), handler.CreateCategory)
	app.Put("/category/:id",middleware.Permission("kelola category"), handler.UpdateCategory)
	app.Delete("/category/:id",middleware.Permission("kelola category"), handler.DeleteCategory)

	// Product Route
	app.Get("/product", handler.GetProduct)
	app.Get("/product/:id", handler.FindProduct)
	app.Post("/product/store", handler.CreateProduct)
	app.Put("/product/:id", handler.UpdateProduct)
	app.Delete("/product/:id", handler.DeleteProduct)
}
