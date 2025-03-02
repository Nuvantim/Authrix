package middleware

import (
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
)


func AuthAndRefreshMiddleware(c *fiber.Ctx) error {
	// Ambil Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}

	// Format header: Bearer <token>
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if tokenString == authHeader { // Jika tidak ada 'Bearer ' di depan, maka token tidak ada
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header must be Bearer <token>",
		})
	}

	// Verifikasi token dengan kunci publik
	publicKey, err := utils.LoadPublicKey()
	if err != nil {
		log.Println("Error loading public key:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error loading public key",
		})
	}

	// Parsing dan verifikasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan signing method sesuai (RS512 dalam kasus ini)
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return publicKey, nil
	})

	if err != nil || !token.Valid {
		log.Println("Error parsing or invalid token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Menyimpan informasi token di context untuk digunakan di handler selanjutnya
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid claims",
		})
	}

	// Set data user dari token ke context Fiber, bisa digunakan di route handler selanjutnya
	c.Locals("user_id", claims["user_id"])
	c.Locals("email", claims["email"])
	c.Locals("roles", claims["roles"])

	// Lanjutkan ke handler berikutnya
	return c.Next()
}
