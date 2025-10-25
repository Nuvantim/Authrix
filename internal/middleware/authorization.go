package middleware

import (
	repo "api/internal/app/repository"
	"api/pkgs/utils/responses"

	"github.com/gofiber/fiber/v2"
)

func Role(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get claims from the context that has been created by jwt
		claims := c.Locals("roles").([]repo.AllRoleClientRow)

		// Check Wheter the user has the required role
		hasRole := false
		for _, role := range claims {
			if role.Name == requiredRole {
				hasRole = true
				break
			}
		}

		if !hasRole {
			return c.Status(fiber.StatusForbidden).JSON(response.Error("authorization", "role forbidden"))
		}

		return c.Next()
	}
}

func Permission(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get claims from the context that has been created by jwt
		claims := c.Locals("roles").([]repo.AllRoleClientRow)

		// Check Wheter the user has the necessary permissions
		hasPermission := false
		for _, role := range claims {
			if perms, ok := role.Permissions.([]interface{}); ok {
				for _, p := range perms {
					if name, ok := p.(map[string]interface{})["name"].(string); ok {
						if name == requiredPermission {
							hasPermission = true
							break
						}
					}
				}
				if hasPermission {
					break
				}
			}
		}

		if !hasPermission {
			return c.Status(fiber.StatusForbidden).JSON(response.Error("authorization", "permission forbidden"))
		}

		return c.Next()
	}
}
