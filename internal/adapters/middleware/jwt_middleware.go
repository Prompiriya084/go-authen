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
		fmt.Println(tokenString)
		jwtToken, err := m.service.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		claims, err := m.service.GetClaims(jwtToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		// fmt.Println(claims)
		//mt.Println(claims["email"])
		userIdfloat, ok := claims["user_id"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid user format")
		}
		// userId := uint(userIdfloat)

		// user, err := userServive.GetUser(userId)
		// if err != nil {
		// 	return c.Status(fiber.StatusUnauthorized).SendString("user not found.")
		// }
		c.Locals("user_id", userIdfloat)
		// c.Locals("role", user.Role)

		fmt.Println(c.Locals("user_id"))
		// fmt.Println(c.Locals("role"))

		return c.Next()
	}

}
