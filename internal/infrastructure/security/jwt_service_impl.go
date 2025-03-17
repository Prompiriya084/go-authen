package security

import (
	"errors"
	"fmt"
	"os"
	"time"

	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Interfaces"
	"github.com/golang-jwt/jwt/v5"
)

type jwtServiceImpl struct {
	secretKey string
}

func NewJwtService() services.IJwtService {
	return &jwtServiceImpl{secretKey: os.Getenv("Jwt_Secret")}
}
func (s *jwtServiceImpl) GenerateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     userId,
		"expiredDate": time.Now().Add(time.Hour * 1).Unix(),
	})
	t, err := token.SignedString([]byte(os.Getenv("Jwt_Secret")))
	if err != nil {
		return "", err
	}
	return t, nil
}
func (s *jwtServiceImpl) ValidateToken(tokenString string) (*jwt.Token, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method.")
		}
		return []byte(s.secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("Invalid token or expired token.")
	}

	// Extract expiration time
	exp, ok := claims["expiredDate"].(float64) // JWT stores `exp` as float64
	if !ok {
		//log.Println("Invalid expiration format")
		return nil, errors.New("Invalid expiration format")
	}

	expirationTime := time.Unix(int64(exp), 0)
	if time.Now().After(expirationTime) {
		//log.Println("Token has expired")
		return nil, errors.New("Token has expired")
	}

	fmt.Println("Token is valid and not expired")
	return token, nil
}
func (s *jwtServiceImpl) GetClaims(token *jwt.Token) (map[string]interface{}, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}

// CheckRole checks if the user has the specified role in the JWT claims.
func (s *jwtServiceImpl) CheckRole(claims map[string]interface{}, role string) bool {
	if roles, ok := claims["roles"].([]interface{}); ok {
		for _, r := range roles {
			if rStr, ok := r.(string); ok && rStr == role {
				return true
			}
		}
	}
	return false
	//return errors.New("Forbidden: Insufficient role")
}
