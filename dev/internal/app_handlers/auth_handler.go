package handler

import (
	"api/internal/app_request"
	"api/internal/app_service"
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

func Login(c *fiber.Ctx) error {
	var login request.Login
	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	user_login, err := service.Login(login)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "sukses", "access_token": user_login})
}
func ResetPassword(c *fiber.Ctx) error {
	var pass request.ResetPassword
	if err := c.BodyParser(&pass); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	update_password, err := service.ResetPassword(pass)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": update_password})
}

// func Logout(c *fiber.Ctx){}
