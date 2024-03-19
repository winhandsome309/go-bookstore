package http

import (
	"go-bookstore/internal/shipping/repository"
	"go-bookstore/internal/shipping/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(c *gin.Engine, db *gorm.DB) {
	shippingRepo := repository.NewShippingRepository(db)
	shippingService := service.NewShippingService(shippingRepo)
	shippingHandlers := NewShippingHandlers(shippingService)

	// shippingRoutes := c.Group("/checkout", middleware.JWTAuth())
	shippingRoutes := c.Group("/checkout")
	{
		shippingRoutes.GET("", shippingHandlers.GetShipping)
		shippingRoutes.GET("/:shippingId", shippingHandlers.GetShippingById)
		shippingRoutes.POST("", shippingHandlers.Checkout)
	}
}
