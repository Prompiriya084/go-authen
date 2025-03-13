package ports

import (
	entities "github.com/Prompiriya084/go-authen/internal/core/entities"
)

type IUserRepository interface {
	GetAll() ([]entities.User, error)
	GetById(id uint) (*entities.User, error)
	Create(user *entities.User) error
	GetWithUserAuthByEmail(email string) (*entities.User, error)
}
