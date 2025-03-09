package services

import (
	"github.com/Prompiriya084/go-authen/internal/core/entities"
	ports "github.com/Prompiriya084/go-authen/internal/core/ports"
	services "github.com/Prompiriya084/go-authen/internal/core/services/interfaces"
	//"github.com/Prompiriya084/go-authen/internal/core/services"
)

type userServiceImpl struct {
	repo ports.IUserRepository
}

func NewUserService(repo ports.IUserRepository) services.IUserService {
	return &userServiceImpl{repo: repo}
}
func (s *userServiceImpl) GetAll() ([]entities.User, error) {
	return s.repo.FindAll()
}
func (s *userServiceImpl) Create(user *entities.User) error {
	return s.repo.Create(user)
}
func (s *userServiceImpl) GetWithUserAuthByEmail(email string) (*entities.User, error) {
	return s.repo.GetWithUserAuthByEmail(email)
}
