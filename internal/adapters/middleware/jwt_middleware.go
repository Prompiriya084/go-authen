package middleware

import (
	infrastructure "github.com/Prompiriya084/go-authen/internal/infrastructure/security"
	"github.com/gofiber/fiber/v3"
)

func JwtMiddleware(requiredRole string) fiber.Handler {
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
			return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
		}

		jwtService := infrastructure.NewJwtService()
		jwtToken, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		claims, err := jwtService.GetClaims(jwtToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		// fmt.Println(claims)
		// fmt.Println(claims["email"])
		if requiredRole != "" {
			if !jwtService.CheckRole(claims, requiredRole) {
				return c.Status(fiber.StatusUnauthorized).SendString("Forbidden: Insufficient role")
			}
		}

		return c.Next()
	}

}
