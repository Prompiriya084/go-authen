package repositories

import (
	entities "github.com/Prompiriya084/go-authen/internal/core/entities"
	ports "github.com/Prompiriya084/go-authen/internal/core/ports"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.IUserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) FindAll() ([]entities.User, error) {
	var users []entities.User
	if result := r.DB.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
func (r *userRepositoryImpl) Create(user *entities.User) error {
	if result := r.DB.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
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
