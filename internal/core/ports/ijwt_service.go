package ports

import "github.com/golang-jwt/jwt/v5"

type IJwtService interface {
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetClaims(token *jwt.Token) (map[string]interface{}, error)
}
