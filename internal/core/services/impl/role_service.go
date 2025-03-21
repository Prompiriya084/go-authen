package services

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Interfaces"
)

type roleServiceImpl struct {
	repo ports.IRoleRepository
}

func NewRoleService(repo ports.IRoleRepository) services.IRoleService {
	return &roleServiceImpl{repo: repo}
}
func (s *roleServiceImpl) GetRoleAll() ([]entities.Role, error) {
	return s.repo.GetRoleAll()
}
func (s *roleServiceImpl) GetRole(id uint) (*entities.Role, error) {
	return s.repo.GetRole(id)
}
func (s *roleServiceImpl) CreateRole(role *entities.Role) error {
	return s.repo.CreateRole(role)
}
