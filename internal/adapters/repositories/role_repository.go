package repositories

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	*GenericRepositoryImpl[entities.Role]
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) ports.IRoleRepository {
	return &roleRepositoryImpl{
		GenericRepositoryImpl: NewGenericRepository[entities.Role](db),
		db:                    db,
	}
}

// func (r *roleRepositoryImpl) GetRoleAll() ([]entities.Role, error) {
// 	var roles []entities.Role
// 	if result := r.DB.Find(&roles); result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return roles, nil

// }

// func (r *roleRepositoryImpl) GetRolesWithFilters(filters *entities.Role, preload []string) (*entities.Role, error) {
// 	var selectedRoles *entities.Role
// 	query := r.DB
// 	for _, p := range preload {
// 		query = query.Preload(p)
// 	}

// 	if filters != nil {
// 		query = query.Where(filters)
// 	}

// 	if result := query.First(&selectedRoles); result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return selectedRoles, nil

// }
// func (r *roleRepositoryImpl) CreateRole(role *entities.Role) error {
// 	if result := r.DB.Create(&role); result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
// func (r *roleRepositoryImpl) UpdateRole(role *entities.Role) error {
// 	if result := r.DB.Save(&role); result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
// func (r *roleRepositoryImpl) DeleteRole(id uint) error {
// 	role := new(entities.Role)
// 	if result := r.DB.Delete(&role, id); result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
