package services

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
)

type IUserService interface {
	GetUserAll() ([]entities.User, error)
	GetUser(id uint) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CreateUser(user *entities.User) error
}
