package services

import (
	"github.com/Prompiriya084/go-authen/internal/core/entities"
)

type IAuthService interface {
	SignIn(userauth *entities.UserAuth) (string, error)
	Register(user *entities.User) error
}
