package handler

import (
	"api/internal/request"
	"api/internal/service"
	"github.com/gofiber/fiber/v2"
)

func GetPermission(c fiber.Ctx) error {
	id := c.Params("id")
	permission, err := service.GetPermission(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(permission)
}

func ListPermission(c fiber.Ctx) error {
	permission, err := service.ListPermission()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(permission)
}

func CreatePermission(c fiber.Ctx) error {
	var data request.Permission
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	permission, err := service.CreatePermission(data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(permission)
}

func UpdatePermission(c fiber.Ctx) error {
	var data request.Permission
	id := c.Params("id")
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	permission, err := service.UpdatePermission(data, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(permission)
}
func DeletePermission(c fiber.Ctx) error {
	id := c.Params("id")
	message, err := service.DeletePermission(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": message})
}
