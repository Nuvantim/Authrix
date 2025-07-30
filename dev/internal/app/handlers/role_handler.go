package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	var data request.Role
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	role, err := service.CreateRole(data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(role)
}
func GetRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	role, err := service.GetRole(int32(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(role)
}
func ListRole(c *fiber.Ctx) error {
	role, err := service.ListRole()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(role)
}
func UpdateRole(c *fiber.Ctx) error {
	var data request.Role
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	role, err := service.UpdateRole(data, int32(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(role)

}
func DeleteRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	msg, err := service.DeleteRole(int32(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": msg})
}
