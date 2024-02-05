// This file contain logic business and ensure to fully synthesis data to send to handler file
package service

import (
	"go-bookstore/internal/product/model"
	"go-bookstore/internal/product/repository"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type IProductService interface {
	GetAllProduct(c *gin.Context) (*[]model.Product, error)
	GetProductById(c *gin.Context, productId string) (*model.Product, error)
	CreateProduct(c *gin.Context, product *model.Product) error
	UpdateProduct(c *gin.Context, productId string, req *model.UpdateProductReq) (*model.Product, error)
	DeleteProduct(c *gin.Context, productId string) error
}

type ProductService struct {
	repo repository.IProductRepo
}

func NewProductService(repo repository.IProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProduct(c *gin.Context) (*[]model.Product, error) {
	products, err := s.repo.GetAllProduct(c)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductById(c *gin.Context, productId string) (*model.Product, error) {
	product, err := s.repo.GetProductById(c, productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) CreateProduct(c *gin.Context, product *model.Product) error {
	err := s.repo.CreateProduct(c, product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateProduct(c *gin.Context, productId string, req *model.UpdateProductReq) (*model.Product, error) {
	product, err := s.repo.GetProductById(c, productId)
	if err != nil {
		log.Error("Product (" + productId + ") not found")
		return nil, err
	}
	err = s.repo.UpdateProduct(c, product, req)
	if err != nil {
		log.Error("Update failed")
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(c *gin.Context, productId string) error {
	product, err := s.repo.GetProductById(c, productId)
	if err != nil {
		return err
	}
	err = s.repo.DeleteProduct(c, product)
	return err
}
