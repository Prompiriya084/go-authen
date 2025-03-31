package ports

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type IUserAuthRepository interface {
	GetUserAuthAll() ([]entities.UserAuth, error)
	GetUserAuthWithFilters(filters *entities.UserAuth, preload []string) (*entities.UserAuth, error)
}
