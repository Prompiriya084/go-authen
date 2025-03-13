package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		fmt.Println(role)
		if !ok {
			return c.Status(fiber.StatusForbidden).SendString("role not found")
		}
		if role != requiredRole {
			return c.Status(fiber.StatusForbidden).SendString("access denied")
		}
		return c.Next()
	}
}
