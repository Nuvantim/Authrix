package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	resp "api/pkgs/utils"
	"github.com/gofiber/fiber/v2"
)

func GetClient(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	client, err := service.GetClient(int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error(err.Error(), "Get Client data"))
	}
	return c.Status(200).JSON(resp.Pass("Get Client Data", client))
}

func ListClient(c *fiber.Ctx) error {
	client, err := service.ListClient()
	if err != nil {
		return c.Status(500).JSON(resp.Error(err.Error(), "List Client data"))
	}
	return c.Status(200).JSON(resp.Pass("Get Client Data", client))
}

func UpdateClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error(err.Error(), "Parser ID"))
	}

	var data request.UpdateClient

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(resp.Error(err.Error(), "Parser JSON"))
	}
	client, err := service.UpdateClient(int32(id), data)
	if err != nil {
		return c.Status(500).JSON(resp.Error(err.Error(), "Update Client"))
	}
	return c.Status(200).JSON(resp.Pass("Update Client data", client))

}

func DeleteClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error(err.Error(), "Parser ID"))
	}

	message, err := service.DeleteClient(int32(id))
	if err != nil {
		c.Status(500).JSON(resp.Error(err.Error(), "Delete Client"))
	}

	return c.Status(200).JSON(resp.Pass(message, struct{}{}))
}
