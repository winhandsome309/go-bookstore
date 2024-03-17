package service

import (
	"testing"

	"go-bookstore/internal/product/model"
	"go-bookstore/internal/product/repository"
	"go-bookstore/internal/product/repository/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProductService_GetAllProduct(t *testing.T) {
	scenarios := []struct {
		name     string
		repo     func(ctrl *gomock.Controller) repository.IProductRepo
		expProds *[]model.Product
		expErr   error
	}{
		{
			name: "Successful case",
			repo: func(ctrl *gomock.Controller) repository.IProductRepo {
				mockProductRepo := mocks.NewMockIProductRepo(ctrl)
				mockProductRepo.EXPECT().GetAllProduct(&gin.Context{}).Return(nil, nil)
				return mockProductRepo
			},
			expProds: nil,
			expErr:   nil,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := NewProductService(scenario.repo(ctrl))
			_, err := s.GetAllProduct(&gin.Context{})
			assert.Equal(t, scenario.expErr, err)
		})
	}
}

func TestProductService_DeleteProduct(t *testing.T) {
	scenarios := []struct {
		name      string
		productId int
		repo      func(ctrl *gomock.Controller) repository.IProductRepo
		expErr    error
	}{
		{
			name:      "successful case",
			productId: 1,
			repo: func(ctrl *gomock.Controller) repository.IProductRepo {
				mockProductRepo := mocks.NewMockIProductRepo(ctrl)
				mockProductRepo.EXPECT().GetProductById(&gin.Context{}, 1).Return(nil, nil)
				mockProductRepo.EXPECT().DeleteProduct(&gin.Context{}, nil).Return(nil)
				return mockProductRepo
			},
			expErr: nil,
		},
	}
	for _, scenario := range scenarios {
		ctrl := gomock.NewController(t)
		s := NewProductService(scenario.repo(ctrl))
		err := s.DeleteProduct(&gin.Context{}, scenario.productId)
		assert.Equal(t, scenario.expErr, err)
	}
}
