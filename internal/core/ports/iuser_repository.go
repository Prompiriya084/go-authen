package ports

import (
	entities "github.com/Prompiriya084/go-authen/internal/core/entities"
)

type IUserRepository interface {
	FindAll() ([]entities.User, error)
	Create(user *entities.User) error
	GetWithUserAuthByEmail(email string) (*entities.User, error)
}
