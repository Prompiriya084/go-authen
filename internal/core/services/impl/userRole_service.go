package services

import (
	"errors"

	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Interfaces"
)

type userRoleServiceImpl struct {
	repo ports.IUserRoleRepository
}

func NewUserRoleService(repo ports.IUserRoleRepository) services.UserRoleService {
	return &userRoleServiceImpl{repo: repo}
}
func (s *userRoleServiceImpl) GetUserRoleAll() ([]entities.UserRole, error) {
	return s.repo.GetUserRoleAll()
}
func (s *userRoleServiceImpl) GetUserRoles(userId uint) ([]entities.UserRole, error) {
	userRoles, err := s.repo.GetUserRoles(userId)
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}
func (s *userRoleServiceImpl) CreateUserRole(userRole *entities.UserRole) error {
	if existingUserRoles, _ := s.GetUserRoles(userRole.UserID); existingUserRoles != nil {
		for _, existingUserRole := range existingUserRoles {
			if existingUserRole.RoleID == userRole.RoleID {
				return errors.New("This user has role exist.")
			}
		}
	}
	if result := s.CreateUserRole(userRole); result.Error != nil {
		return result
	}
	return nil
}
