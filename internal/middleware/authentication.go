package middleware

import (
	"api/pkgs/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
)

func BearerAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tokenString string
		authHeader := c.Get("Authorization")
		authCookie := c.Get("Set-Cookie")

		// Ambil token dari header Authorization
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// Validasi access token
		if tokenString != "" {
			token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
				// Pastikan metode signing adalah RS512
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || token.Method.Alg() != "RS256" {
					return nil, jwt.ErrSignatureInvalid
				}
				return utils.PublicKey, nil
			})

			// Jika access token valid, set user context dan lanjutkan
			if err == nil && token.Valid {
				if claims, ok := token.Claims.(*utils.Claims); ok {
					c.Locals("user_id", claims.UserID)
					c.Locals("email", claims.Email)
					c.Locals("roles", claims.Roles)
					c.Set("Authorization", authHeader)
					return c.Next()
				}
			} else {
				// Jika access token tidak valid, coba refresh token
				if authHeader != "" && authCookie != "" {
					refreshToken, err := jwt.ParseWithClaims(authCookie, &utils.RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
						// Pastikan metode signing adalah RS512
						if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || token.Method.Alg() != "RS256" {
							return nil, jwt.ErrSignatureInvalid
						}
						return utils.PublicKey, nil
					})

					if err == nil && refreshToken.Valid {
						if claims, ok := refreshToken.Claims.(*utils.RefreshClaims); ok {
							newAccessToken, err := utils.AutoRefreshToken(claims.UserID)
							if err == nil {
								// Validasi token baru
								token, err := jwt.ParseWithClaims(newAccessToken, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
									if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || token.Method.Alg() != "RS256" {
										return nil, jwt.ErrSignatureInvalid
									}
									return utils.PublicKey, nil
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
									log.Printf("error validating new access token: %v", err)
								}
							} else {
								log.Printf("error refreshing access token: %v", err)
							}
						}
					} else {
						log.Printf("refresh token invalid: %v", err)
					}
				}
			}
		}

		// Jika kedua token tidak valid, kembalikan response unauthorized
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Error("authentication", "authentication required"))
	}
}
