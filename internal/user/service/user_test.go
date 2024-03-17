package service

import (
	"errors"
	"go-bookstore/internal/user/model"
	"go-bookstore/internal/user/repository"
	"go-bookstore/internal/user/repository/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Register(t *testing.T) {
	g := &gin.Context{}
	scenarios := []struct {
		name    string
		userReq *model.UserReq
		repo    func(ctrl *gomock.Controller) repository.IUserRepository
		expErr  error
	}{
		{
			name:    "Success case",
			userReq: &model.UserReq{},
			repo: func(ctrl *gomock.Controller) repository.IUserRepository {
				mockUserRepository := mocks.NewMockIUserRepository(ctrl)
				mockUserRepository.EXPECT().CreateUser(gomock.Any(), &model.User{
					Id: 1,
				}).Return(nil)
				return mockUserRepository
			},
			expErr: nil,
		},
		{
			name: "Error case",
			userReq: &model.UserReq{
				Email: "thang@gmail.com",
			},
			repo: func(ctrl *gomock.Controller) repository.IUserRepository {
				mockUserRepository := mocks.NewMockIUserRepository(ctrl)
				mockUserRepository.EXPECT().CreateUser(gomock.Any(), &model.User{
					Email: "thang@gmail.com",
				}).Return(errors.New("connection error"))
				return mockUserRepository
			},
			expErr: errors.New("connection error"),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			u := &UserService{repo: scenario.repo(ctrl)}
			err := u.Register(g, scenario.userReq)
			assert.Equal(t, scenario.expErr, err)
		})
	}
}

func TestUserService_SignIn(t *testing.T) {
	g := &gin.Context{}
	scenarios := []struct {
		name      string
		userLogin *model.UserLogin
		repo      func(ctrl *gomock.Controller) repository.IUserRepository
		expErr    []interface{}
	}{
		{
			name: "Success case",
			userLogin: &model.UserLogin{
				Email:    "thang@gmail.com",
				Password: "30092002",
			},
			repo: func(ctrl *gomock.Controller) repository.IUserRepository {
				mockUserRepository := mocks.NewMockIUserRepository(ctrl)
				mockUserRepository.EXPECT().GetUserByEmail(g, "thang@gmail.com").Return(&model.User{
					Email:    "thang@gmail.com",
					Password: "$2a$14$uGxmZ1lMXKcIQYSt/gsI3.YQudMK/xcJl3A0MISCj3gCh4jCcgb62",
				}, nil)

				return mockUserRepository
			},
			expErr: []interface{}{"abc", "abc", nil},
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			u := &UserService{repo: scenario.repo(ctrl)}
			var err []interface{}
			a, b, c := u.SignIn(g, scenario.userLogin)
			err = append(err, a)
			err = append(err, b)
			err = append(err, c)
			assert.Equal(t, scenario.expErr, err)
		})
	}
}
