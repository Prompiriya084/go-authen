package middleware

import (
	"fmt"
	"strings"

	services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type RoleMiddleware struct {
	service services.UserRoleService
}

func NewRoleMiddleware(service services.UserRoleService) *RoleMiddleware {
	return &RoleMiddleware{service: service}
}
func (m *RoleMiddleware) RequiredRole(requiredRole string) fiber.Handler {
	return func(c fiber.Ctx) error {
		fmt.Println("Required Role method!!!!")
		localsUserId := c.Locals("user_id").(uuid.UUID)
		// fmt.Println(localsUserId)
		// userId, ok := localsUserId.(uint)
		// fmt.Println(userId)
		// if !ok {
		// 	return c.Status(fiber.StatusForbidden).SendString("role not found")
		// }
		// if userId != requiredRole {
		// 	return c.Status(fiber.StatusForbidden).SendString("access denied")
		// }

		userRoles, err := m.service.GetUserRolesById(localsUserId)
		if err != nil {
			return c.Status(fiber.StatusForbidden).SendString("Cannot fetch roles")
		}
		fmt.Println(userRoles[0])
		for _, userRole := range userRoles {
			fmt.Println(userRole.Role.Name)
			if strings.EqualFold(userRole.Role.Name, requiredRole) {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).SendString("Permission denied.")

	}
}
