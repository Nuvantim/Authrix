package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	resp "api/pkgs/utils"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	var id = c.Locals("user_id").(int32)
	if id == 0 {
		return c.Status(401).JSON(resp.Error("get user_id", "unauthorized"))
	}
	// Get Account by id
	user, err := service.GetProfile(id)
	if err != nil {
		return c.Status(500).JSON(resp.Error("get profile", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass("get profile", user))
}

func UpdateAccount(c *fiber.Ctx) error {
	var id = c.Locals("user_id").(int32)
	if id == 0 {
		return c.Status(401).JSON(resp.Error("get user_id", "unauthorized"))
	}
	var user request.UpdateAccount
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(resp.Error("parser json", err.Error()))
	}

	// validate data
	if err := resp.Validates(user); err != nil {
		return c.Status(400).JSON(resp.Error("validation data", err.Error()))
	}

	userUpdate, err := service.UpdateAccount(user, id)
	if err != nil {
		return c.Status(500).JSON(resp.Error("update account", err.Error()))
	}

	return c.Status(200).JSON(resp.Pass("update account", userUpdate))
}

func DeleteAccount(c *fiber.Ctx) error {
	var id = c.Locals("user_id").(int32)
	if id == 0 {
		return c.Status(401).JSON(resp.Error("get user_id", "unauthorized"))
	}
	msg, err := service.DeleteAccount(id)
	if err != nil {
		return c.Status(500).JSON(resp.Error("delete account", err.Error()))
	}
	return c.Status(200).JSON(resp.Pass(msg, struct{}{}))
}
