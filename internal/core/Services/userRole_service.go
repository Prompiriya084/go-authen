package services

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	"github.com/google/uuid"
)

type UserRoleService interface {
	GetUserRoleAll() ([]entities.UserRole, error)
	GetUserRolesById(id uuid.UUID) ([]entities.UserRole, error)
	CreateUserRole(userRole *entities.UserRole) error
}
