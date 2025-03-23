package ports

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type IUserRoleRepository interface {
	GetUserRoleAll() ([]entities.UserRole, error)
	GetUserRolesByStruct(userRole *entities.UserRole) ([]entities.UserRole, error)
	CreateUserRole(userRole *entities.UserRole) error
}
