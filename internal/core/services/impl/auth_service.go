package services

import (
	"errors"
	"fmt"

	request "github.com/Prompiriya084/go-authen/Internal/Adapters/Request"
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Repositories"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Interfaces"
	"github.com/google/uuid"
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
	user, err := s.repoUser.GetWithUserAuthByEmail(userAuth.Email)
	if err != nil {
		return "", errors.New("email or password is incorrect.")
	}
	fmt.Println(user)
	hashedpassword := user.UserAuth.Password

	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedpassword),
		[]byte(userAuth.Password),
	); err != nil {
		return "", errors.New("email or password is incorrect.")
	}

	token, err := s.jwtService.GenerateToken(user.UserAuth.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authServiceImpl) Register(requestRegister *request.RequestRegister) error {

	if user, _ := s.repoUser.GetWithUserAuthByEmail(requestRegister.Email); user != nil {
		return errors.New("This user is already exists.")
	}
	if role, _ := s.repoRole.GetRole(requestRegister.Role); role == nil {
		return errors.New("Role not found.")
	}
	user := entities.User{
		ID:      uuid.New(),
		Name:    requestRegister.Name,
		Surname: requestRegister.Surname,
	}
	fmt.Println(user)
	//Check userRole is exists ?
	if userRoles, _ := s.repoUserRole.GetUserRolesByStruct(&entities.UserRole{UserID: user.ID}); userRoles != nil {
		for _, userRole := range userRoles {
			if userRole.Role.ID == requestRegister.Role {
				return errors.New("This user has this role exist.")
			}
		}
	}
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.UserAuth.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.UserAuth.ID = user.ID
	user.UserAuth.Email = requestRegister.Email
	user.UserAuth.Password = string(hashedpassword)
	fmt.Println(user)
	if err := s.repoUser.CreateUser(&user); err != nil {
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
