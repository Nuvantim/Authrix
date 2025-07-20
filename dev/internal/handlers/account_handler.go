package handler

import (
	"api/internal/request"
	"api/internal/service"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	// id := c.Locals("id")
	var userID int32 = 2
	if userID == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}
	// Get Account by id
	user, err := service.GetProfile(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(user)
}
func UpdateAccount(c *fiber.Ctx) error {
	// id := c.Locals("id")
	var userID int32 = 2
	if userID == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}
	var user request.UpdateAccount
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	userUpdate, err := service.UpdateAccount(user, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).JSON(userUpdate)
}

func DeleteAccount(c *fiber.Ctx) error {
	// id := c.Locals("id")
	var userID int32 = 2
	if userID == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}
	msg, err := service.DeleteAccount(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": msg})
}
