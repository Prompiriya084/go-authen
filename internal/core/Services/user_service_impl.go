package services

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	//"github.com/Prompiriya084/go-authen/internal/core/services"
)

type userServiceImpl struct {
	repo ports.IUserRepository
}

func NewUserService(repo ports.IUserRepository) IUserService {
	return &userServiceImpl{repo: repo}
}
func (s *userServiceImpl) GetUserAll() ([]entities.User, error) {
	return s.repo.GetUserAll()
}
func (s *userServiceImpl) GetUser(id uint) (*entities.User, error) {
	filters := map[string]interface{}{
		"id": id,
	}
	return s.repo.GetUserWithFilters(filters, []string{"UserAuth"})
}
func (s *userServiceImpl) CreateUser(user *entities.User) error {
	return s.repo.CreateUser(user)
}
func (s *userServiceImpl) GetUserWithUserAuthByEmail(email string) (*entities.User, error) {
	return s.repo.GetUserWithFilters(map[string]interface{}{
		"email": email,
	}, []string{"UserAuth"})
}
