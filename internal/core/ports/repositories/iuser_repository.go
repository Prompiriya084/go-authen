package ports

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	"github.com/google/uuid"
)

type IUserRepository interface {
	GetUserAll() ([]entities.User, error)
	GetUserWithFilters(filters map[string]interface{}, preload []string) (*entities.User, error)
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User) error
	DeleteUser(id uuid.UUID) error
}
