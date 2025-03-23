package middleware

import (
	"fmt"

	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Interfaces"
	"github.com/gofiber/fiber/v3"
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

		userId, ok := claims["user_id"].(uint)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid user format")
		}
		// c.Locals("user_id", userIdfloat)

		// fmt.Println(c.Locals("user_id"))
		fiber.Locals[uint](c, "user_id", userId)
		test := fiber.Locals[uint](c, "user_id")
		fmt.Println(test)

		return c.Next()

	}

}
