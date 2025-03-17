package services

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Interfaces"
	//"github.com/Prompiriya084/go-authen/internal/core/services"
)

type userServiceImpl struct {
	repo ports.IUserRepository
}

func NewUserService(repo ports.IUserRepository) services.IUserService {
	return &userServiceImpl{repo: repo}
}
func (s *userServiceImpl) GetUserAll() ([]entities.User, error) {
	return s.repo.GetUserAll()
}
func (s *userServiceImpl) GetUser(id uint) (*entities.User, error) {
	return s.repo.GetUser(id)
}
func (s *userServiceImpl) CreateUser(user *entities.User) error {
	return s.repo.CreateUser(user)
}
func (s *userServiceImpl) GetWithUserAuthByEmail(email string) (*entities.User, error) {
	return s.repo.GetWithUserAuthByEmail(email)
}
