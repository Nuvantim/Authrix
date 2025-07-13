package server

import (
	"github.com/gofiber/fiber/v2"
	"api/internal/routes"
	"api/config"
)

func ServerGo() *fiber.App{
	app := fiber.New()
	routes.Setup(app)
	return app

}
