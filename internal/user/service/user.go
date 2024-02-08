package service

import (
	"go-bookstore/internal/user/model"
	"go-bookstore/internal/user/repository"
	"go-bookstore/pkg/utils"
	"time"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("30092002")

type IUserService interface {
	Register(c *gin.Context, userReq *model.UserReq) error
	SignIn(c *gin.Context, userLogin *model.UserLogin) (string, time.Time, error)
	RefreshToken(c *gin.Context) error
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{repo: repo}

}

func (s *UserService) Register(c *gin.Context, userReq *model.UserReq) error {
	var userNew model.User
	utils.Merge(&userNew, userReq)
	err := s.repo.CreateUser(c, &userNew)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) SignIn(c *gin.Context, userLogin *model.UserLogin) (string, time.Time, error) {
	user, err := s.repo.GetUserByEmail(c, userLogin.Email)
	if err != nil {
		return "", time.Time{}, err
	}
	if user.Password != userLogin.Password {
		return "", time.Time{}, errors.New("Wrong password")
	}
	expTime := time.Now().Add(5 * time.Minute)
	tokenContent := jwt.MapClaims{
		"payload": map[string]interface{}{
			"id":    user.Id,
			"email": user.Email,
			"role":  user.Role,
		},
		"exp": expTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenContent)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Time{}, errors.New("Failed to generate access token")
	}
	return tokenString, expTime, nil
}

func (s *UserService) RefreshToken(c *gin.Context) error {
	return nil

}
