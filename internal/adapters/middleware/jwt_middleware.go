package middleware

import (
	"fmt"

	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Interfaces"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type JwtMiddleware struct {
	service services.IJwtService
}

func NewJwtMiddleware(jwtService services.IJwtService) *JwtMiddleware {
	return &JwtMiddleware{service: jwtService}
}
func (m *JwtMiddleware) AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		fmt.Println("Jwt middleware excuted!!!!")
		// tokenString := c.Get("Authorization") // Get token from header
		// if bearerString := strings.Split(tokenString, " "); bearerString[0] != "Bearer" {
		// 	fmt.Println(bearerString[0])
		// 	// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		// 	return c.Status(fiber.StatusUnauthorized).SendString("Invalid token format")
		// }
		// // Remove "Bearer " prefix if present
		// if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		// 	tokenString = tokenString[7:]
		// }

		tokenString := c.Cookies("jwt") //follow by create cookies name in sign in service
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Missing token.")
		}
		jwtToken, err := m.service.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		claims, err := m.service.GetClaims(jwtToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		userIdStr, ok := claims["user_id"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid user format")
		}
		userId, err := uuid.Parse(userIdStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid UUID format")
		}

		c.Locals("user_id", userId)

		// fmt.Println(c.Locals("user_id"))

		return c.Next()

	}

}
