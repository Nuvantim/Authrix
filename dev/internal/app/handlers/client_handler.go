package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	"github.com/gofiber/fiber/v2"
)

func GetClient(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	client, err := service.GetClient(int32(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(client)
}

func ListClient(c *fiber.Ctx) error {
	data, err := service.ListClient()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(data)
}

func UpdateClient(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var data request.UpdateClient

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	client,role err := service.UpdateClient(int32(id), data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON("client":client,"role":role)

}

func DeleteClient(c *fiber.Ctx) error {
	return nil
}
