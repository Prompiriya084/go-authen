package ports

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type IUserAuthRepository interface {
	IRepository[entities.UserAuth]
	// GetUserAuthAll() ([]entities.UserAuth, error)
	// GetUserAuthWithFilters(filters *entities.UserAuth, preload []string) (*entities.UserAuth, error)
}
