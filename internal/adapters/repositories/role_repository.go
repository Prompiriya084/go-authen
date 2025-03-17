package repositories

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) ports.IRoleRepository {
	return &roleRepositoryImpl{DB: db}
}
func (r *roleRepositoryImpl) GetRoleAll() ([]entities.Role, error) {
	var roles []entities.Role
	if result := r.DB.Find(&roles); result.Error != nil {
		return nil, result.Error
	}

	return roles, nil

}
func (r *roleRepositoryImpl) GetRole(id uint) (*entities.Role, error) {
	var role entities.Role
	if result := r.DB.First(&role, id); result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}
func (r *roleRepositoryImpl) CreateRole(role *entities.Role) error {
	if result := r.DB.Create(&role); result.Error != nil {
		return result.Error
	}
	return nil
}
