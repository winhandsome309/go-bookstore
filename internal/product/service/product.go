package service

import (
	"go-bookstore/internal/product/model"
	"go-bookstore/internal/product/repository"

	"github.com/gin-gonic/gin"
)

type ProductService struct {
	repo *repository.ProductRepo
}

func NewProductService(repo *repository.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProduct(c *gin.Context) (*[]model.Product, error) {
	products, err := s.repo.GetAllProduct(c)
	if err != nil {
		return nil, err
	}
	return products, nil
}
