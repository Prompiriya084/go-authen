package services

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
)

type roleServiceImpl struct {
	repo ports.IRoleRepository
}

func NewRoleService(repo ports.IRoleRepository) IRoleService {
	return &roleServiceImpl{repo: repo}
}
func (s *roleServiceImpl) GetRoleAll() ([]entities.Role, error) {
	return s.repo.GetAll(nil, nil)
}
func (s *roleServiceImpl) GetRole(id uint) (*entities.Role, error) {
	return s.repo.Get(&entities.Role{
		ID: id,
	}, nil)
}
func (s *roleServiceImpl) CreateRole(role *entities.Role) error {
	return s.repo.Add(role)
}
func (s *roleServiceImpl) UpdateRole(role *entities.Role) error {
	selectedRole, err := s.repo.Get(&entities.Role{
		ID: role.ID,
	}, nil)
	if err != nil {
		return err
	}
	selectedRole.Name = role.Name
	if err := s.repo.Update(selectedRole); err != nil {
		return err
	}

	return nil
}
func (s *roleServiceImpl) DeleteRole(id uint) error {
	selectedRole, err := s.repo.Get(&entities.Role{
		ID: id,
	}, nil)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(selectedRole); err != nil {
		return err
	}

	return nil
}
