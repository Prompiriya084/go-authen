package middleware

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware(c fiber.Ctx) error {
	// authHeader := c.Get("Authorization") // Get Authorization header
	// if authHeader == "" {
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }
	// // Ensure it starts with "Bearer "
	// tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	// if tokenString == authHeader { // Token not prefixed with "Bearer "
	// 	return c.SendStatus(fiber.StatusUnauthorized) //.JSON(fiber.Map{"error": "Invalid token format"})
	// }

	// token, err := jwt.ParseWithClaims(tokenString, jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(os.Getenv("Jwt_Secret")), nil
	// })
	// if err != nil || !token.Valid {
	// 	fmt.Print(err.Error())
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }
	tokenString := c.Get("Authorization") // Get token from header
	if bearerString := strings.Split(tokenString, " "); bearerString[0] != "Bearer" {
		fmt.Println(bearerString[0])
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	// Remove "Bearer " prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse Token
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	return []byte(os.Getenv("Jwt_Secret")), nil
	// })
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("Jwt_Secret")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Extract expiration time
	exp, ok := claims["expiredDate"].(float64) // JWT stores `exp` as float64
	if !ok {
		//log.Println("Invalid expiration format")
		return errors.New("Invalid expiration format")
	}

	expirationTime := time.Unix(int64(exp), 0)
	if time.Now().After(expirationTime) {
		//log.Println("Token has expired")
		return errors.New("Token has expired")
	}

	fmt.Println("Token is valid and not expired")
	return c.Next()
}
