package services

import (
	"errors"

	"github.com/Prompiriya084/go-authen/internal/core/entities"
	ports "github.com/Prompiriya084/go-authen/internal/core/ports/repositories"
	services "github.com/Prompiriya084/go-authen/internal/core/services/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repo       ports.IUserRepository
	jwtService services.IJwtService
}

func NewAuthService(repo ports.IUserRepository, jwtService services.IJwtService) services.IAuthService {
	return &AuthServiceImpl{repo: repo, jwtService: jwtService}
}
func (s *AuthServiceImpl) SignIn(userAuth *entities.UserAuth) (string, error) {
	selectedUserAuth, err := s.repo.GetWithUserAuthByEmail(userAuth.Email)
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

	token, err := s.jwtService.GenerateToken(int(selectedUserAuth.ID))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthServiceImpl) Register(user *entities.User) error {

	if user, _ := s.repo.GetWithUserAuthByEmail(user.UserAuth.Email); user != nil {
		return errors.New("This user is already exists.")
	}
	user.Role = "user"
	//user.Role = "user"
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.UserAuth.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.UserAuth.Password = string(hashedpassword)

	if err := s.repo.Create(user); err != nil {
		return err
	}
	return nil
}
