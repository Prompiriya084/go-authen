package services

import (
	"errors"

	request "github.com/Prompiriya084/go-authen/Internal/Adapters/Request"
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"

	// services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	repoUser     ports.IUserRepository
	repoUserAuth ports.IUserAuthRepository
	repoUserRole ports.IUserRoleRepository
	repoRole     ports.IRoleRepository
	jwtService   IJwtService
}

func NewAuthService(repoUser ports.IUserRepository,
	repoUserAuth ports.IUserAuthRepository,
	repoUserRole ports.IUserRoleRepository,
	repoRole ports.IRoleRepository,
	jwtService IJwtService) IAuthService {
	return &authServiceImpl{
		repoUser:     repoUser,
		repoUserAuth: repoUserAuth,
		repoUserRole: repoUserRole,
		repoRole:     repoRole,
		jwtService:   jwtService,
	}
}
func (s *authServiceImpl) SignIn(userAuth *entities.UserAuth) (string, error) {
	selectedUserAuth, err := s.repoUserAuth.Get(&entities.UserAuth{
		Email: userAuth.Email,
	}, nil)
	if err != nil {
		return "", errors.New("email or password is incorrect.")
	}
	hashedpassword := selectedUserAuth.Password

	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedpassword),
		[]byte(userAuth.Password),
	); err != nil {
		return "", errors.New("email or password is incorrect.")
	}

	token, err := s.jwtService.GenerateToken(selectedUserAuth.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authServiceImpl) Register(requestRegister *request.RequestRegister) error {

	if userAuth, _ := s.repoUserAuth.Get(&entities.UserAuth{
		Email: requestRegister.Email,
	}, nil); userAuth != nil {
		return errors.New("This user is exists.")
	}
	if role, _ := s.repoRole.Get(&entities.Role{
		ID: requestRegister.Role,
	}, nil); role == nil {
		return errors.New("Role not found.")
	}
	user := entities.User{
		ID:      uuid.New(),
		Name:    requestRegister.Name,
		Surname: requestRegister.Surname,
	}
	//Check userRole is exists ?
	if userRoles, _ := s.repoUserRole.GetUserRolesWithFilters(&entities.UserRole{
		UserID: user.ID,
	}, nil); userRoles != nil {
		for _, userRole := range userRoles {
			if userRole.Role.ID == requestRegister.Role {
				return errors.New("This user has this role exist.")
			}
		}
	}
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(requestRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.UserAuth.ID = user.ID
	user.UserAuth.Email = requestRegister.Email
	user.UserAuth.Password = string(hashedpassword)

	if err := s.repoUser.Add(&user); err != nil {
		return err
	}

	var userRole entities.UserRole
	userRole.UserID = user.ID
	userRole.RoleID = requestRegister.Role
	if err := s.repoUserRole.CreateUserRole(&userRole); err != nil {
		return err
	}
	return nil
}
