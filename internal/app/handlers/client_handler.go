package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	resp "api/pkgs/utils"
	"github.com/gofiber/fiber/v2"
)

func GetClient(c *fiber.Ctx) error {
	// Get id
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("get id", err.Error()))
	}
	// ID Validation
	id, err := resp.ValID(params)
	if err != nil {
		return c.Status(500).JSON(resp.Error("validation", err.Error()))
	}
	client, err := service.GetClient(id)
	if err != nil {
		return c.Status(500).JSON(resp.Error("get client data", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("get client data", client))
}

func ListClient(c *fiber.Ctx) error {
	client, err := service.ListClient()
	if err != nil {
		return c.Status(500).JSON(resp.Error("list client", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("list client", client))
}

func UpdateClient(c *fiber.Ctx) error {
	// Get id
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("get id", err.Error()))
	}
	// ID Validation
	id, err := resp.ValID(params)
	if err != nil {
		return c.Status(500).JSON(resp.Error("validation", err.Error()))
	}
	var data request.UpdateClient

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(resp.Error("parser json", err.Error()))
	}
	// validate data
	if err := resp.Validates(data); err != nil {
		return c.Status(400).JSON(resp.Error("validation data", err.Error()))
	}

	client, err := service.UpdateClient(id, data)
	if err != nil {
		return c.Status(500).JSON(resp.Error("update client", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("update client", client))
}

func DeleteClient(c *fiber.Ctx) error {
	// Get id
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("get id", err.Error()))
	}
	// ID Validation
	id, err := resp.ValID(params)
	if err != nil {
		return c.Status(500).JSON(resp.Error("validation", err.Error()))
	}
	message, err := service.DeleteClient(id)
	if err != nil {
		c.Status(500).JSON(resp.Error("delete client", err.Error()))
	}

	return c.Status(200).JSON(resp.Pass(message, struct{}{}))
}
