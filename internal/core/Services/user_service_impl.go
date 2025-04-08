package services

import (
	"encoding/binary"
	"fmt"

	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	"github.com/google/uuid"
	//"github.com/Prompiriya084/go-authen/internal/core/services"
)

type userServiceImpl struct {
	repo ports.IUserRepository
}

func NewUserService(repo ports.IUserRepository) IUserService {
	return &userServiceImpl{repo: repo}
}
func (s *userServiceImpl) GetUserAll() ([]entities.User, error) {
	return s.repo.GetAll(nil, nil)
}
func (s *userServiceImpl) GetUser(id uint) (*entities.User, error) {
	// filters := map[string]interface{}{
	// 	"id": id,
	// }
	var uuid uuid.UUID
	binary.BigEndian.PutUint64(uuid[8:], uint64(id))
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
