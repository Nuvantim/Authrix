package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	var data request.Role
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(data); err != nil {
		return c.Status(400).JSON(response.Error("validation data", err.Error()))
	}

	role, err := service.CreateRole(data)
	if err != nil {
		return c.Status(500).JSON(response.Error("create role", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("create role", role))
}
func GetRole(c *fiber.Ctx) error {
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
	role, err := service.GetRole(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("get role", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("get role", role))
}
func ListRole(c *fiber.Ctx) error {
	role, err := service.ListRole()
	if err != nil {
		return c.Status(500).JSON(response.Error("list role", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("list role", role))
}
func UpdateRole(c *fiber.Ctx) error {
	var data request.Role
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
	//
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(data); err != nil {
		return c.Status(400).JSON(response.Error("validation data", err.Error()))
	}

	role, err := service.UpdateRole(data, id)
	if err != nil {
		return c.Status(500).JSON(response.Error("update password", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("update role", role))

}
func DeleteRole(c *fiber.Ctx) error {
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
	msg, err := service.DeleteRole(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("delete role", err.Error()))
	}
	return c.Status(200).JSON(response.Pass(msg, struct{}{}))
}
