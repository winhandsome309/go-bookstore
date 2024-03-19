package http

import (
	orderRepository "go-bookstore/internal/order/repository"
	orderService "go-bookstore/internal/order/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(c *gin.Engine, db *gorm.DB) {
	orderReposity := orderRepository.NewOrderRepository(db)
	orderService := orderService.NewOrderService(orderReposity)
	paymentHandler := NewPaymentHandler(orderService)
	c.POST("/payment/:id", paymentHandler.Payment)
}
