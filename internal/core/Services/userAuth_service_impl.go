package services

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
)

type userAuthServiceImpl struct {
	repo ports.IUserAuthRepository
}

func NewUserAuthService(repo ports.IUserAuthRepository) UserAuthService {
	return &userAuthServiceImpl{repo: repo}
}

func (s *userAuthServiceImpl) GetUserAuthAll() ([]entities.UserAuth, error) {
	return s.repo.GetAll(nil, nil)
}
func (s *userAuthServiceImpl) GetUserAuthByEmail(email string) (*entities.UserAuth, error) {
	return s.repo.Get(&entities.UserAuth{
		Email: email,
	}, nil)
}
