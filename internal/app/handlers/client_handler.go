package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	resp "api/pkgs/utils"
	"github.com/gofiber/fiber/v2"
	"math"
)

func GetClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil{
		return c.Status(400).JSON(resp.Error("get id", err.Error()))
	}
	// validate range int32
	if id < math.MinInt32 || id > math.MaxInt32{
		return c.Status(400).JSON(resp.Error("validation", "out of int32 range"))
	}
	client, err := service.GetClient(int32(id))
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
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("parser id", err.Error()))
	}
	// validate range int32
	if id < math.MinInt32 || id > math.MaxInt32{
		return c.Status(400).JSON(resp.Error("validation", "out of int32 range"))
	}
	var data request.UpdateClient

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(resp.Error("parser json", err.Error()))
	}
	// validate data
	if err := resp.Validates(data); err != nil {
		return c.Status(400).JSON(resp.Error("validation data", err.Error()))
	}

	client, err := service.UpdateClient(int32(id), data)
	if err != nil {
		return c.Status(500).JSON(resp.Error("update client", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("update client", client))
}

func DeleteClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("parser id", err.Error()))
	}
	// validate range int32
	if id < math.MinInt32 || id > math.MaxInt32{
		return c.Status(400).JSON(resp.Error("validation", "out of int32 range"))
	}
	message, err := service.DeleteClient(int32(id))
	if err != nil {
		c.Status(500).JSON(resp.Error("delete client", err.Error()))
	}

	return c.Status(200).JSON(resp.Pass(message, struct{}{}))
}
