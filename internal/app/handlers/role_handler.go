package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	resp "api/pkgs/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	var data request.Role
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(resp.Error("parser json", err.Error()))
	}

	// validate data
	if err := resp.Validates(data); err != nil {
		return c.Status(400).JSON(resp.Error("validation data", err.Error()))
	}

	role, err := service.CreateRole(data)
	if err != nil {
		return c.Status(500).JSON(resp.Error("create role", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("create role", role))
}
func GetRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("get id", err.Error()))
	}
	role, err := service.GetRole(int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error("get role", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("get role", role))
}
func ListRole(c *fiber.Ctx) error {
	role, err := service.ListRole()
	if err != nil {
		return c.Status(500).JSON(resp.Error("list role", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("list role", role))
}
func UpdateRole(c *fiber.Ctx) error {
	var data request.Role
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("get id", err.Error()))
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(resp.Error("parser json", err.Error()))
	}

	// validate data
	if err := resp.Validates(data); err != nil {
		return c.Status(400).JSON(resp.Error("validation data", err.Error()))
	}

	role, err := service.UpdateRole(data, int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error("update password", err.Error()))
	}

	return c.Status(200).JSON(resp.Pass("update role", role))

}
func DeleteRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("get id", err.Error()))
	}
	msg, err := service.DeleteRole(int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error("delete role", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass(msg, struct{}{}))
}
