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

func (r *userRoleRepositoryImpl) GetUserRolesByStruct(userRole *entities.UserRole) ([]entities.UserRole, error) {
	var selectedUserRoles []entities.UserRole
	// if result := r.DB.Preload("User").Preload("Role").
	// 	Where("user_id = ?", userId).
	// 	Find(&userRoles); result.Error != nil {
	// 	return nil, result.Error
	// }
	if result := r.DB.Where(&userRole).Find(&selectedUserRoles); result.Error != nil {
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
