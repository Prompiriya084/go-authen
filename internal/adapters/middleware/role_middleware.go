package middleware

import (
	"fmt"
	"strings"

	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Interfaces"
	"github.com/gofiber/fiber/v3"
)

type RoleMiddleware struct {
	service services.UserRoleService
}

func NewRoleMiddleware(service services.UserRoleService) *RoleMiddleware {
	return &RoleMiddleware{service: service}
}
func (m *RoleMiddleware) RequiredRole(requiredRole string) fiber.Handler {
	return func(c fiber.Ctx) error {
		userId, ok := c.Locals("user_id").(uint)
		fmt.Println(userId)
		if !ok {
			return c.Status(fiber.StatusForbidden).SendString("role not found")
		}
		// if userId != requiredRole {
		// 	return c.Status(fiber.StatusForbidden).SendString("access denied")
		// }

		roles, err := m.service.GetUserRoles(userId)
		if err != nil {
			return c.Status(fiber.StatusForbidden).SendString("Cannot fetch roles")
		}

		for _, role := range roles {
			if strings.EqualFold(role.Role.Name, requiredRole) {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).SendString("Permission denied.")
	}
}
