package ports

import (
	entities "github.com/Prompiriya084/go-authen/internal/core/entities"
)

type IUserRepository interface {
	FindAll() ([]entities.User, error)
	Create(user *entities.User) error
	GetUserWithUserAuthByEmail(email string) (*entities.User, error)
}
