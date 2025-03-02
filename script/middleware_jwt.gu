package middleware

import (
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"strings"
)

var (
	jwtSecret     = []byte(os.Getenv("API_KEY"))
	refreshSecret = []byte(os.Getenv("REFRESH_KEY"))
)

func AuthAndRefreshMiddleware(c *fiber.Ctx) error {
	var tokenString string
	authHeader := c.Get("Authorization")
	authCookie := c.Cookies("refresh_token")

	// Retrieve the token from the Authorization header or cookies
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
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
		} else {
			if authHeader != "" && authCookie != "" {
				refreshToken, err := jwt.ParseWithClaims(authCookie, &utils.RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
					return refreshSecret, nil
				})

				if err == nil && refreshToken.Valid {
					if claims, ok := refreshToken.Claims.(*utils.RefreshClaims); ok {
						newAccessToken, err := utils.AutoRefressToken(claims.UserID)
						if err == nil {
							// Validate the new access token
							token, err := jwt.ParseWithClaims(newAccessToken, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
								return jwtSecret, nil
							})
							if err == nil && token.Valid {
								if claims, ok := token.Claims.(*utils.Claims); ok {
									c.Locals("user_id", claims.UserID)
									c.Locals("email", claims.Email)
									c.Locals("roles", claims.Roles)
									c.Set("Authorization", "Bearer "+newAccessToken)
									return c.Next()
								}
							} else {
								log.Printf("Error validating new access token: %v", err) // Log the error
							}
						} else {
							log.Printf("Error refreshing access token: %v", err) // Log the error
						}
					}
				} else {
					log.Printf("Refresh token invalid: %v", err) // Log the error
				}
			}
		}
	}

	// Get Refresh token from cookie

	// If both tokens are invalid, return an unauthorized response
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Authentication required",
	})
}
