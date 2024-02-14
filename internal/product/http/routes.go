package http

import (
	"go-bookstore/internal/product/repository"
	"go-bookstore/internal/product/service"

	"go-bookstore/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := NewProductHandler(productService)

	productRoute := r.Group("/products", middleware.JWTAuth())
	{
		productRoute.GET("", productHandler.GetAllProduct)
		productRoute.POST("", productHandler.CreateProduct)
		productRoute.GET("/:id", productHandler.GetProductById)
		productRoute.PATCH("/:id", productHandler.UpdateProduct)
		productRoute.DELETE("/:id", productHandler.DeleteProduct)
	}
}
