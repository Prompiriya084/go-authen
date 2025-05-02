package ports_repositories

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type IRoleRepository interface {
	IRepository[entities.Role]
	// GetRoleAll() ([]entities.Role, error)
	// GetRolesWithFilters(filters *entities.Role, preload []string) (*entities.Role, error)
	// CreateRole(role *entities.Role) error
	// UpdateRole(role *entities.Role) error
	// DeleteRole(id uint) error
}
