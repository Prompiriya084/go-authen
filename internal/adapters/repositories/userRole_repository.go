package repositories

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	"gorm.io/gorm"
)

type userRoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) ports.IUserRoleRepository {
	return &userRoleRepositoryImpl{DB: db}
}
func (r *userRoleRepositoryImpl) GetUserRoleAll() ([]entities.UserRole, error) {
	var userRoles []entities.UserRole

	if result := r.DB.Find(&userRoles); result.Error != nil {
		return nil, result.Error
	}
	return userRoles, nil
}
func (r *userRoleRepositoryImpl) GetUserRolesWithFilters(filters *entities.UserRole, preload []string) ([]entities.UserRole, error) {
	var selectedUserRoles []entities.UserRole
	query := r.DB
	for _, p := range preload {
		query = query.Preload(p)
	}

	if filters != nil {
		query = query.Where(filters)
	}

	if result := query.Find(&selectedUserRoles); result.Error != nil {
		return nil, result.Error
	}

	return selectedUserRoles, nil
}
func (r *userRoleRepositoryImpl) CreateUserRole(userRole *entities.UserRole) error {
	if result := r.DB.Create(&userRole); result.Error != nil {
		return result.Error
	}
	return nil
}
