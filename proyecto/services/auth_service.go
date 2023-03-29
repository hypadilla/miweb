package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"proyecto/models"
	"proyecto/repositories"
)

type AuthService interface {
	Login(email, password string) (string, error)
	Register(user *models.User) error
}

type authService struct {
	userRepo     repositories.UserRepository
	secretKey    string
	tokenExpires int64
}

func NewAuthService(userRepo repositories.UserRepository, secretKey string, tokenExpires int64) AuthService {
	return &authService{userRepo: userRepo, secretKey: secretKey, tokenExpires: tokenExpires}
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	/*if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password 1")
	}*/

	if password != user.Password {
		return "", errors.New("invalid email or password")
	}

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}

func (s *authService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}

	user.Password = string(hashedPassword)

	return s.userRepo.CreateUser(user)
}
