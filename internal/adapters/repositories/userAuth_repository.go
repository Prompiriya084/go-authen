package repositories

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	"gorm.io/gorm"
)

type userAuthRepositoryImpl struct {
	*GenericRepositoryImpl[entities.UserAuth]
	db *gorm.DB
}

func NewUserAuthRepository(db *gorm.DB) ports.IUserAuthRepository {
	return &userAuthRepositoryImpl{
		GenericRepositoryImpl: NewGenericRepository[entities.UserAuth](db),
		db:                    db,
	}
}

// func (r *userAuthRepositoryImpl) GetUserAuthAll() ([]entities.UserAuth, error) {
// 	var userAuths []entities.UserAuth
// 	if result := r.db.Find(&userAuths); result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return userAuths, nil
// }
// func (r *userAuthRepositoryImpl) GetUserAuthWithFilters(filters *entities.UserAuth, preload []string) (*entities.UserAuth, error) {
// 	var selectedUserAuth *entities.UserAuth
// 	// fmt.Println("UserAuthRepo : ", userAuth)
// 	// if result := r.db.Where(&userAuth).First(&selectedUserAuth); result.Error != nil {
// 	// 	return nil, result.Error
// 	// }
// 	// return selectedUserAuth, nil

// 	query := r.db
// 	for _, p := range preload {
// 		query = query.Preload(p)
// 	}

// 	if filters != nil {
// 		query = query.Where(filters)
// 	}

// 	if result := query.First(&selectedUserAuth); result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return selectedUserAuth, nil
// }
