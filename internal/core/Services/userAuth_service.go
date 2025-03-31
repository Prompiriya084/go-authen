package services

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type UserAuthService interface {
	GetUserAuthAll() ([]entities.UserAuth, error)
	GetUserAuthByEmail(email string) (*entities.UserAuth, error)
}
