package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	resp "api/pkgs/utils"
	"github.com/gofiber/fiber/v2"
	"math"
)

func GetPermission(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("get id", err.Error()))
	}
	// validate range int32
	if id < math.MinInt32 || id > math.MaxInt32{
		return c.Status(400).JSON(resp.Error("validation", "out of int32 range"))
	}
	permission, err := service.GetPermission(int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error("get permission", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("get permission", permission))
}

func ListPermission(c *fiber.Ctx) error {
	permission, err := service.ListPermission()
	if err != nil {
		return c.Status(500).JSON(resp.Error("list permission", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("list permission", permission))
}

func CreatePermission(c *fiber.Ctx) error {
	var data request.Permission
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(resp.Error("parser json", err.Error()))
	}
	// validate data
	if err := resp.Validates(data); err != nil {
		return c.Status(400).JSON(resp.Error("validation data", err.Error()))
	}
	permission, err := service.CreatePermission(data)
	if err != nil {
		return c.Status(500).JSON(resp.Error("create permission", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("create permission", permission))
}

func UpdatePermission(c *fiber.Ctx) error {
	var data request.Permission
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("get id", err.Error()))
	}
	// validate range int32
	if id < math.MinInt32 || id > math.MaxInt32{
		return c.Status(400).JSON(resp.Error("validation", "out of int32 range"))
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(resp.Error("parser json", err.Error()))
	}
	// validation data
	if err := resp.Validates(data); err != nil {
		return c.Status(400).JSON(resp.Error("validation data", err.Error()))
	}
	permission, err := service.UpdatePermission(data, int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error("update permission", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("update permission", permission))
}
func DeletePermission(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("get id", err.Error()))
	}
	// validate range int32
	if id < math.MinInt32 || id > math.MaxInt32{
		return c.Status(400).JSON(resp.Error("validation", "out of int32 range"))
	}
	message, err := service.DeletePermission(int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error("delete permission", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass(message, struct{}{}))
}
