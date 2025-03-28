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
func (r *userRepositoryImpl) GetUser(userId uint) (*entities.User, error) {
	var user entities.User
	if result := r.DB.First(&user, userId); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
func (r *userRepositoryImpl) GetByStruct(user *entities.User) (*entities.User, error) {
	var selectedUser *entities.User
	if result := r.DB.Where(&user).First(&selectedUser); result.Error != nil {
		return nil, result.Error
	}
	return selectedUser, nil
}
func (r *userRepositoryImpl) GetWithUserAuthByEmail(email string) (*entities.User, error) {
	// var userAuth entities.UserAuth
	// if err := db.Where("email = ?", email).First(&userAuth).Error; err != nil {
	// 	fmt.Println("UserAuth not found")
	// 	return nil, err
	// }

	// var user entities.User
	// if err := db.Where("user_auth_id = ?", userAuth.ID).Preload("UserAuth").First(&user).Error; err != nil {
	// 	fmt.Println("User not found")
	// 	return nil, err
	// }
	var user entities.User
	result := r.DB.Preload("UserAuth").Where("user_auths.email = ?", email).Joins("JOIN user_auths ON user_auths.id = users.user_auth_id").First(&user)
	if result.Error != nil {
		//log.Fatalf("Error get book: %v", result.Error)
		return nil, result.Error
	}

	return &user, nil
}
func (r *userRepositoryImpl) CreateUser(user *entities.User) error {
	if result := r.DB.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
