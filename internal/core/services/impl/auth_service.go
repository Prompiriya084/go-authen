package services

import (
	"errors"

	request "github.com/Prompiriya084/go-authen/Internal/Adapters/Request"
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Interfaces"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	repoUser     ports.IUserRepository
	repoUserRole ports.IUserRoleRepository
	repoRole     ports.IRoleRepository
	jwtService   services.IJwtService
}

func NewAuthService(repoUser ports.IUserRepository,
	repoUserRole ports.IUserRoleRepository,
	repoRole ports.IRoleRepository,
	jwtService services.IJwtService) services.IAuthService {
	return &authServiceImpl{
		repoUser:     repoUser,
		repoUserRole: repoUserRole,
		repoRole:     repoRole,
		jwtService:   jwtService,
	}
}
func (s *authServiceImpl) SignIn(userAuth *entities.UserAuth) (string, error) {
	selectedUserAuth, err := s.repoUser.GetWithUserAuthByEmail(userAuth.Email)
	if err != nil {
		return "", errors.New("email or password is incorrect.")
	}
	hashedpassword := selectedUserAuth.UserAuth.Password

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

func (s *authServiceImpl) Register(request *request.RequestRegister) error {
	var user entities.User
	user = request.User

	if user, _ := s.repoUser.GetWithUserAuthByEmail(user.UserAuth.Email); user != nil {
		return errors.New("This user is already exists.")
	}

	var userRole entities.UserRole
	userRole.UserID = user.ID
	userRole.RoleID = request.Role

	if role, _ := s.repoRole.GetRole(request.Role); role != nil {
		return errors.New("Role not found.")
	}
	//Check userRole is exists ?
	if roles, _ := s.repoUserRole.GetUserRoles(userRole.UserID); roles != nil {
		for _, role := range roles {
			if role.Role.ID == request.Role {
				return errors.New("This user has this role exist.")
			}
		}
	}
	// user.Role = "user"
	//user.Role = "user"
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.UserAuth.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.UserAuth.Password = string(hashedpassword)

	if err := s.repoUser.CreateUser(&user); err != nil {
		return err
	}
	if err := s.repoUserRole.CreateUserRole(&userRole); err != nil {
		return err
	}
	return nil
}
