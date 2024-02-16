package service

import (
	"go-bookstore/internal/user/model"
	"go-bookstore/internal/user/repository"
	"go-bookstore/pkg/utils"
	"time"

	"errors"

	"go-bookstore/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// var jwtKey = []byte("30092002")

type IUserService interface {
	Register(c *gin.Context, userReq *model.UserReq) error
	SignIn(c *gin.Context, userLogin *model.UserLogin) (string, int64, error)
	SignOut(c *gin.Context) (string, error)
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

func (s *UserService) SignIn(c *gin.Context, userLogin *model.UserLogin) (string, int64, error) {
	cfg := config.GetConfig()
	user, err := s.repo.GetUserByEmail(c, userLogin.Email)
	if err != nil {
		return "", 0, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		return "", 0, errors.New("Wrong password")
	}
	expTime := time.Now().Add(20 * time.Minute).Unix()
	tokenContent := jwt.MapClaims{
		"payload": map[string]interface{}{
			"id":    user.Id,
			"email": user.Email,
			"role":  user.Role,
		},
		"exp": expTime,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenContent)
	tokenString, err := token.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		return "", 0, errors.New("Failed to generate access token")
	}
	return tokenString, expTime, nil
}

func (s *UserService) SignOut(c *gin.Context) (string, error) {
	cfg := config.GetConfig()
	tokenContent := jwt.MapClaims{
		"payload": "",
		"exp":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenContent)
	tokenString, err := token.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		return "", errors.New("Failed to generate access token")
	}
	return tokenString, nil
}
