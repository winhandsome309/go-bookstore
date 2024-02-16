package http

import (
	"go-bookstore/internal/shipping/model"
	"go-bookstore/internal/shipping/service"

	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ShippingHandlers struct {
	service service.IShippingService
}

func NewShippingHandlers(service service.IShippingService) *ShippingHandlers {
	return &ShippingHandlers{service: service}
}

func (h *ShippingHandlers) GetShipping(c *gin.Context) {
	shippings, err := h.service.GetAllShipping(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "get fail",
		})
		return
	}
	c.JSON(http.StatusOK, shippings)
}

func (h *ShippingHandlers) GetShippingById(c *gin.Context) {
	shippingId := c.Param("shippingId")
	shipping, err := h.service.GetShippingById(c, shippingId)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "get fail",
		})
		return
	}
	c.JSON(http.StatusOK, shipping)
}

func (h *ShippingHandlers) Checkout(c *gin.Context) {
	var shipping model.Shipping
	if err := c.ShouldBind(&shipping); c.Request.Body == nil || err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "get fail",
		})
		return
	}
	err := h.service.Checkout(c, &shipping)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "get fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Checkout successfully",
	})
}
