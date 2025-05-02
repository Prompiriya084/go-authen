package ports_repositories

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	"github.com/google/uuid"
)

type IUserRoleRepository interface {
	GetUserRoleAll() ([]entities.UserRole, error)
	GetUserRolesWithFilters(filters *entities.UserRole, preload []string) ([]entities.UserRole, error)
	CreateUserRole(userRole *entities.UserRole) error
	UpdateUserRole(userRole *entities.UserRole) error
	DeleteUserRole(id uuid.UUID) error
}
