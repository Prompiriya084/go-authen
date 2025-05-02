package ports_repositories

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type IUserRepository interface {
	IRepository[entities.User]
	// GetWithFilters(filters map[string]interface{}, preload []string) (*entities.User, error)
	// GetUserAll() ([]entities.User, error)
	// GetUserAllWithFilters(filters map[string]interface{}, preload []string) ([]entities.User, error)
	// GetUserWithFilters(filters map[string]interface{}, preload []string) (*entities.User, error)
	// CreateUser(user *entities.User) error
	// UpdateUser(user *entities.User) error
	// DeleteUser(id uuid.UUID) error
}
