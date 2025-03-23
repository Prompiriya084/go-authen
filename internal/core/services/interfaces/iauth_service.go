package services

import (
	request "github.com/Prompiriya084/go-authen/Internal/Adapters/Request"
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
)

type IAuthService interface {
	SignIn(userauth *entities.UserAuth) (string, error)
	Register(requestRegister *request.RequestRegister) error
}
