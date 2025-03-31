package repositories

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.IUserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) GetUserAll() ([]entities.User, error) {
	var users []entities.User
	if result := r.DB.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// func (r *userRepositoryImpl) GetUser(userId uint) (*entities.User, error) {
// 	var user entities.User
// 	if result := r.DB.First(&user, userId); result.Error != nil {
// 		return nil, result.Error
// 	}

//		return &user, nil
//	}
func (r *userRepositoryImpl) GetUserWithFilters(filters map[string]interface{}, preload []string) (*entities.User, error) {
	query := r.DB
	var selectedUser *entities.User
	for _, p := range preload {
		query = query.Preload(p)
	}

	if len(filters) > 0 {
		query = query.Where(filters)
	}

	if result := query.First(&selectedUser); result.Error != nil {
		return nil, result.Error
	}
	return selectedUser, nil
}

func (r *userRepositoryImpl) CreateUser(user *entities.User) error {
	if result := r.DB.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
