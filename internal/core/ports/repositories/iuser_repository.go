package ports

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
)

type IUserRepository interface {
	GetUserAll() ([]entities.User, error)
	GetUser(userId uint) (*entities.User, error)
	GetByStruct(user *entities.User) (*entities.User, error)
	GetWithUserAuthByEmail(email string) (*entities.User, error)
	CreateUser(user *entities.User) error
}
