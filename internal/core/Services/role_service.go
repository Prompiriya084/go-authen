package services

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type IRoleService interface {
	GetRoleAll() ([]entities.Role, error)
	GetRole(id uint) (*entities.Role, error)
	CreateRole(role *entities.Role) error
	UpdateRole(role *entities.Role) error
	DeleteRole(id uint) error
}
