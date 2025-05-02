package services

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports_repositories "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
)

type userAuthServiceImpl struct {
	repo ports_repositories.IUserAuthRepository
}

func NewUserAuthService(repo ports_repositories.IUserAuthRepository) UserAuthService {
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
