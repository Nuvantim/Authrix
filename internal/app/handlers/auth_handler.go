package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/request"
	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
)

func SendOTP(c *fiber.Ctx) error {
	var otp request.OtpToken
	if err := c.BodyParser(&otp); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(otp); err != nil {
		return c.Status(422).JSON(response.Error("validation data", err.Error()))
	}

	send, err := service.SendOTP(otp.Email)
	if err != nil {
		return c.Status(500).JSON(response.Error("send otp", err.Error()))
	}

	return c.Status(200).JSON(response.Pass(send, struct{}{}))
}

func Register(c *fiber.Ctx) error {
	var regist request.Register
	if err := c.BodyParser(&regist); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(regist); err != nil {
		return c.Status(422).JSON(response.Error("validation data", err.Error()))
	}

	user_regist, err := service.Register(regist)
	if err != nil {
		return c.Status(500).JSON(response.Error("register account", err.Error()))
	}
	return c.Status(200).JSON(response.Pass(user_regist, struct{}{}))
}

func Login(c *fiber.Ctx) error {
	var login request.Login
	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(login); err != nil {
		return c.Status(422).JSON(response.Error("validation data", err.Error()))
	}

	access, refresh, err := service.Login(login)
	if err != nil {
		return c.Status(500).JSON(response.Error("login account", err.Error()))
	}
	// Set Cookie with refresh token
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refresh,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	})

	// Set Response with access token
	return c.Status(200).JSON(response.Pass("login account", struct {
		Token string `json:"access_token"`
	}{Token: access}))
}
func ResetPassword(c *fiber.Ctx) error {
	var pass request.ResetPassword
	if err := c.BodyParser(&pass); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(pass); err != nil {
		return c.Status(422).JSON(response.Error("validation data", err.Error()))
	}

	update_password, err := service.ResetPassword(pass)
	if err != nil {
		return c.Status(500).JSON(response.Error("reset password", err.Error()))
	}
	return c.Status(200).JSON(response.Pass(update_password, struct{}{}))
}

func Logout(c *fiber.Ctx) error {
	// Clear the access token cookie
	c.Set("Authorization", "")

	// Clear the refresh token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   true, // Set to true in production
		SameSite: "Strict",
		Path:     "/",
		MaxAge:   -1, // Delete cookie
	})

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
