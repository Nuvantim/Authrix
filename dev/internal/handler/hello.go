package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx){
	return c.Status(200).JSON(fiber.Map{
		"message" : "Hello World",
	})
}