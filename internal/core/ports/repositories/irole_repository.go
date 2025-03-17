package ports

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type IRoleRepository interface {
	GetRoleAll() ([]entities.Role, error)
	GetRole(id uint) (*entities.Role, error)
	CreateRole(role *entities.Role) error
}
