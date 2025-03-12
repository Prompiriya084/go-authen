package services

import (
	"errors"
	"os"
	"time"

	"github.com/Prompiriya084/go-authen/internal/core/entities"
	ports "github.com/Prompiriya084/go-authen/internal/core/ports/repositories"
	services "github.com/Prompiriya084/go-authen/internal/core/services/interfaces"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repo ports.IUserRepository
}

func NewAuthService(repo ports.IUserRepository) services.IAuthService {
	return &AuthServiceImpl{repo: repo}
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":       userAuth.Email,
		"expiredDate": time.Now().Add(time.Hour * 1).Unix(),
	})
	t, err := token.SignedString([]byte(os.Getenv("Jwt_Secret")))
	if err != nil {
		return "", err
	}

	return t, nil
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
