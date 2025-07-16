package handler

import (
	"api/internal/request"
	"api/internal/service"
	"github.com/gofiber/fiber/v2"
)

func SendOTP(c *fiber.Ctx) error {
	var otp request.OtpToken
	if err := c.BodyParser(&otp); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	send, err := service.SendOTP(otp.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": send})
}

func Register(c *fiber.Ctx) error {
	var regist request.Register
	if err := c.BodyParser(&regist); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	user_regist, err := service.Register(regist)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": user_regist})
}

// func Login(c *fiber.Ctx) error{}
// func UpdatePassword(c *fiber.Ctx) error{}
// func Logout(c *fiber.Ctx){}
