package ports

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type IUserRoleRepository interface {
	GetUserRoleAll() ([]entities.UserRole, error)
	GetUserRolesWithFilters(filters *entities.UserRole, preload []string) ([]entities.UserRole, error)
	CreateUserRole(userRole *entities.UserRole) error
}
