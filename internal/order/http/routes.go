package http

import (
	"go-bookstore/internal/order/repository"
	"go-bookstore/internal/order/service"

	"go-bookstore/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(c *gin.Engine, db *gorm.DB) {
	orderReposity := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderReposity)
	orderHandler := NewOrderHandler(orderService)

	orderRoute := c.Group("/orders", middleware.JWTAuth())
	{
		orderRoute.GET("", orderHandler.GetOrder)
		orderRoute.POST("", orderHandler.CreateUpdateOrderLine)
		orderRoute.DELETE("/:orderId", orderHandler.DeleteOrder)
	}

	orderLineRoute := c.Group("/orderlines", middleware.JWTAuth())
	{
		orderLineRoute.GET("", orderHandler.GetOrderLines)
		orderLineRoute.DELETE("", orderHandler.DeleteOrderLineById)
		// orderLineRoute.POST("", orderHandler.GetOrderLines)
	}
}
