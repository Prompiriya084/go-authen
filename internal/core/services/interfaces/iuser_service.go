package services

import (
	entities "github.com/Prompiriya084/go-authen/internal/core/entities"
)

type IUserService interface {
	GetAll() ([]entities.User, error)
	Create(user *entities.User) error
	GetWithUserAuthByEmail(email string) (*entities.User, error)
}
