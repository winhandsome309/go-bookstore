package repository

import (
	"go-bookstore/internal/product/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) GetAllProduct(c *gin.Context) (*[]model.Product, error) {
	var products []model.Product
	r.db.Find(&products)
	return &products, nil
}
