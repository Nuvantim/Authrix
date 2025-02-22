package middleware

import (
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

var (
	jwtSecret     = []byte(os.Getenv("API_KEY"))
	refreshSecret = []byte(os.Getenv("REFRESH_KEY"))
)

func AuthAndRefreshMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenString := ""

	// Retrieve the token from the Authorization header or cookies
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		tokenString = c.Cookies("access_token")
	}

	// Try to validate the access token
	if tokenString != "" {
		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// If the access token is valid, set user context and proceed
		if err == nil && token.Valid {
			if claims, ok := token.Claims.(*utils.Claims); ok {
				c.Locals("user_id", claims.UserID)
				c.Locals("email", claims.Email)
				c.Locals("roles", claims.Roles)
				c.Set("Authorization", authHeader)
				return c.Next()
			}
		}
	}

	// If both tokens are invalid, return an unauthorized response
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Authentication required",
	})
}
