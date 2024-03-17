package service

import (
	"errors"
	model_order "go-bookstore/internal/order/model"
	"go-bookstore/internal/shipping/model"
	"go-bookstore/internal/shipping/repository"
	"go-bookstore/internal/shipping/repository/mocks"
	model_user "go-bookstore/internal/user/model"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShipping_Checkout(t *testing.T) {
	shippingTest := model.Shipping{
		Id:               1,
		OrderId:          1,
		FirstName:        "Thang",
		LastName:         "Nguyen Xuan",
		Email:            "thang@gmail.com",
		Phone:            "123456789",
		Address:          "abc",
		ShippingProvince: "01",
		ShippingDistrict: "001",
		ShippingWard:     "0001",
	}
	orderTest := model_order.Order{
		Id:         1,
		CustomerID: 1,
		TotalPrice: 20000,
	}
	userTest := model_user.User{
		Id:      1,
		Email:   "thang@gmail.com",
		Balance: 100000,
	}
	scenarios := []struct {
		name     string
		shipping model.Shipping
		order    model_order.Order
		user     model_user.User
		repo     func(ctrl *gomock.Controller, ship *model.Shipping, order *model_order.Order, user *model_user.User) repository.IShippingRepository
		expErr   func(ship *model.Shipping, order *model_order.Order, user *model_user.User) error
	}{
		{
			name:     "check if user balance is reduced",
			shipping: shippingTest,
			order:    orderTest,
			user:     userTest,
			repo: func(ctrl *gomock.Controller, ship *model.Shipping, order *model_order.Order, user *model_user.User) repository.IShippingRepository {
				shippingRepo := mocks.NewMockIShippingRepository(ctrl)
				shippingRepo.EXPECT().Checkout(&gin.Context{}, ship).Return(nil)
				shippingRepo.EXPECT().GetOrderById(&gin.Context{}, strconv.Itoa(ship.OrderId)).Return(order, nil)
				shippingRepo.EXPECT().GetUserById(&gin.Context{}, strconv.Itoa(order.CustomerID)).Return(user, nil)
				shippingRepo.EXPECT().UpdateUser(&gin.Context{}, user).Return(errors.New("abc"))
				return shippingRepo
			},
			expErr: func(ship *model.Shipping, order *model_order.Order, user *model_user.User) error {
				if user.Balance == userTest.Balance-orderTest.TotalPrice {
					return nil
				}
				return errors.New("Something went wrong")
			},
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			shippingService := NewShippingService(scenario.repo(ctrl, &scenario.shipping, &scenario.order, &scenario.user))
			shippingService.Checkout(&gin.Context{}, &scenario.shipping)
			assert.Equal(t, nil, scenario.expErr(&scenario.shipping, &scenario.order, &scenario.user))
		})
	}
}
