package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	var id = c.Locals("user_id").(int32)
	if id == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
	}
	// Get Account by id
	user, err := service.GetProfile(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(user)
}
func UpdateAccount(c *fiber.Ctx) error {
	var id = c.Locals("user_id").(int32)
	if id == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
	}
	var user request.UpdateAccount
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	userUpdate, err := service.UpdateAccount(user, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).JSON(userUpdate)
}

func DeleteAccount(c *fiber.Ctx) error {
	var id = c.Locals("user_id").(int32)
	if id == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
	}
	msg, err := service.DeleteAccount(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": msg})
}
