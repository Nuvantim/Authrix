package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	"github.com/gofiber/fiber/v2"
)

func GetClient(c *fiber.Ctx) error {
	// Get id
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}
	// ID Validation
	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(500).JSON(response.Error("validation", err.Error()))
	}
	client, err := service.GetClient(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("get client data", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("get client data", client))
}

func ListClient(c *fiber.Ctx) error {
	client, err := service.ListClient()
	if err != nil {
		return c.Status(500).JSON(response.Error("list client", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("list client", client))
}

func UpdateClient(c *fiber.Ctx) error {
	// Get id
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}
	// ID Validation
	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(500).JSON(response.Error("validation", err.Error()))
	}
	var data request.UpdateClient

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(data); err != nil {
		return c.Status(422).JSON(response.Error("validation data", err.Error()))
	}

	client, err := service.UpdateClient(id, data)
	if err != nil {
		return c.Status(500).JSON(response.Error("update client", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("update client", client))
}

func DeleteClient(c *fiber.Ctx) error {
	// Get id
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}
	// ID Validation
	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(500).JSON(response.Error("validation", err.Error()))
	}
	message, err := service.DeleteClient(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("delete client", err.Error()))
	}

	return c.Status(200).JSON(response.Pass(message, struct{}{}))
}
