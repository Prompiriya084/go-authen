package services

import "github.com/golang-jwt/jwt/v5"

type IJwtService interface {
	GenerateToken(userId uint) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetClaims(token *jwt.Token) (map[string]interface{}, error)
	CheckRole(claims map[string]interface{}, role string) bool
}
