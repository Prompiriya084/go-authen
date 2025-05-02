package services

import (
	"errors"

	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports_repositories "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	"github.com/google/uuid"
)

type userRoleServiceImpl struct {
	repo ports_repositories.IUserRoleRepository
}

func NewUserRoleService(repo ports_repositories.IUserRoleRepository) UserRoleService {
	return &userRoleServiceImpl{repo: repo}
}
func (s *userRoleServiceImpl) GetUserRoleAll() ([]entities.UserRole, error) {
	return s.repo.GetUserRoleAll()
}
func (s *userRoleServiceImpl) GetUserRolesById(id uuid.UUID) ([]entities.UserRole, error) {
	preload := []string{"Role"}
	userRoles, err := s.repo.GetUserRolesWithFilters(&entities.UserRole{
		UserID: id,
	}, preload)
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}
func (s *userRoleServiceImpl) CreateUserRole(userRole *entities.UserRole) error {
	if existingUserRoles, _ := s.repo.GetUserRolesWithFilters(userRole, nil); existingUserRoles != nil {
		for _, existingUserRole := range existingUserRoles {
			if existingUserRole.RoleID == userRole.RoleID {
				return errors.New("This user has role exist.")
			}
		}
	}
	if err := s.repo.CreateUserRole(userRole); err != nil {
		return err
	}
	return nil
}
