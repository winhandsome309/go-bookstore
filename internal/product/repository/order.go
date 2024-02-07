// This file to interact with database
package repository

import (
	"go-bookstore/internal/product/model"
	"go-bookstore/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IProductRepo interface {
	GetAllProduct(c *gin.Context) (*[]model.Product, error)
	GetProductById(c *gin.Context, id string) (*model.Product, error)
	CreateProduct(c *gin.Context, product *model.Product) error
	UpdateProduct(c *gin.Context, product *model.Product, req *model.UpdateProductReq) error
	DeleteProduct(c *gin.Context, product *model.Product) error
}

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) GetAllProduct(c *gin.Context) (*[]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (r *ProductRepo) GetProductById(c *gin.Context, id string) (*model.Product, error) {
	var product model.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepo) CreateProduct(c *gin.Context, product *model.Product) error {
	err := r.db.Create(&product).Error
	return err
}

func (r *ProductRepo) UpdateProduct(c *gin.Context, product *model.Product, req *model.UpdateProductReq) error {
	utils.Merge(product, req)
	err := r.db.Save(product).Error
	return err
}

func (r *ProductRepo) DeleteProduct(c *gin.Context, product *model.Product) error {
	err := r.db.Unscoped().Delete(product).Error
	return err
}
