package repository

import (
	"go-bookstore/internal/user/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(c *gin.Context, userReq *model.User) error
	GetUserByEmail(c *gin.Context, email string) (*model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}

}

func (r *UserRepository) CreateUser(c *gin.Context, userReq *model.User) error {
	err := r.db.Create(userReq).Error
	return err
}

func (r *UserRepository) GetUserByEmail(c *gin.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
