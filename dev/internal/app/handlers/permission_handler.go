package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	resp "api/pkgs/utils"
	"github.com/gofiber/fiber/v2"
)

func GetPermission(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error(err.Error(),"get id"))
	}
	permission, err := service.GetPermission(int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error(err.Error(),"get permission"))
	}
	return c.Status(200).JSON(resp.Pass("get permission",permission))
}

func ListPermission(c *fiber.Ctx) error {
	permission, err := service.ListPermission()
	if err != nil {
		return c.Status(500).JSON(resp.Error(err.Error(), "list permission"))
	}
	return c.Status(200).JSON(resp.Pass("list permission",permission))
}

func CreatePermission(c *fiber.Ctx) error {
	var data request.Permission
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(resp.Error(err.Error(), "parser json"))
	}
	// validate data
	if err := resp.Validates(data);err != nil{
		return c.Status(400).JSON(resp.Error(err.Error(),"validation data"))
	}
	permission, err := service.CreatePermission(data)
	if err != nil {
		return c.Status(500).JSON(resp.Error(err.Error(), "create permission"))
	}
	return c.Status(200).JSON(resp.Pass("create permission",permission))
}

func UpdatePermission(c *fiber.Ctx) error {
	var data request.Permission
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error(err.Error(), "get id"))
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(resp.Error(err.Error(), "parser json"))
	}
	// validation data
	if err := resp.Validates(data);err != nil{
		return c.Status(400).JSON(resp.Error(err.Error(),"validation data"))
	}
	permission, err := service.UpdatePermission(data, int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error(err.Error(), "update permission"))
	}
	return c.Status(200).JSON(resp.Pass("update permission",permission))
}
func DeletePermission(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error(err.Error(), "get id"))
	}
	message, err := service.DeletePermission(int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error(err.Error(), "delete permission"))
	}
	return c.Status(200).JSON(resp.Pass(message, struct{}{}))
}
