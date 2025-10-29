package services

import (
	"fmt"

	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports_repositories "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	"github.com/google/uuid"
	//"github.com/Prompiriya084/go-authen/internal/core/services"
)

type userServiceImpl struct {
	repo ports_repositories.IUserRepository
}

func NewUserService(repo ports_repositories.IUserRepository) IUserService {
	return &userServiceImpl{repo: repo}
}
func (s *userServiceImpl) GetUserAll(filters *entities.User) ([]entities.User, error) {
	return s.repo.GetAll(nil, nil)
}
func (s *userServiceImpl) GetUser(id string) (*entities.User, error) {
	// filters := map[string]interface{}{
	// 	"id": id,
	// }
	uuid, _ := uuid.Parse(id)
	return s.repo.Get(&entities.User{
		ID: uuid,
	}, []string{"UserAuth"})
}
func (s *userServiceImpl) GetUserByEmail(email string) (*entities.User, error) {
	users, err := s.repo.GetAll(nil, []string{"UserAuth"})
	if err != nil {
		return nil, err
	}
	var selectedUser *entities.User
	for _, user := range users {
		fmt.Println(user)
		if user.UserAuth.Email == email {
			selectedUser = &user
			break
		}
	}

	return selectedUser, nil
}
func (s *userServiceImpl) CreateUser(user *entities.User) error {
	return s.repo.Add(user)
}
