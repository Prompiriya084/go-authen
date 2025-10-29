package mockitem_repositories

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type MockUserAuthRepository struct {
	*MockRepositoryImpl[entities.UserAuth]
}
