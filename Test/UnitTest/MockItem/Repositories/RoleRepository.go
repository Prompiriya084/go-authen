package mockitem_repositories

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type MockRoleRepository struct {
	*MockRepositoryImpl[entities.Role]
}
