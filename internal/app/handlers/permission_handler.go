package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	"github.com/gofiber/fiber/v2"
)

func GetPermission(c *fiber.Ctx) error {
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

	permission, err := service.GetPermission(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("get permission", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("get permission", permission))
}

func ListPermission(c *fiber.Ctx) error {
	permission, err := service.ListPermission()
	if err != nil {
		return c.Status(500).JSON(response.Error("list permission", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("list permission", permission))
}

func CreatePermission(c *fiber.Ctx) error {
	var data request.Permission
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(data); err != nil {
		return c.Status(400).JSON(response.Error("validation data", err.Error()))
	}

	permission, err := service.CreatePermission(data)
	if err != nil {
		return c.Status(500).JSON(response.Error("create permission", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("create permission", permission))
}

func UpdatePermission(c *fiber.Ctx) error {
	var data request.Permission
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

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}
	// validation data
	if err := validate.BodyStructs(data); err != nil {
		return c.Status(400).JSON(response.Error("validation data", err.Error()))
	}

	permission, err := service.UpdatePermission(data, id)
	if err != nil {
		return c.Status(500).JSON(response.Error("update permission", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("update permission", permission))
}
func DeletePermission(c *fiber.Ctx) error {
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

	message, err := service.DeletePermission(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("delete permission", err.Error()))
	}

	return c.Status(200).JSON(response.Pass(message, struct{}{}))
}
