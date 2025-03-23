package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type IJwtService interface {
	GenerateToken(uuid uuid.UUID) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetClaims(token *jwt.Token) (map[string]interface{}, error)
	CheckRole(claims map[string]interface{}, role string) bool
}
