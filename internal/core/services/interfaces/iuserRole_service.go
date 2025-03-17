package services

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type UserRoleService interface {
	GetUserRoleAll() ([]entities.UserRole, error)
	GetUserRoles(userId uint) ([]entities.UserRole, error)
	CreateUserRole(userRole *entities.UserRole) error
}
