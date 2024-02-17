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

// GetShipping godoc
//
//	@Summary	get shipping info of user
//	@Tags		shipping
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{array}	model.Shipping
//	@Router		/checkout [get]
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

// GetShippingById godoc
//
//	@Summary	get shipping info by id
//	@Tags		shipping
//	@Produce	json
//	@Security	ApiKeyAuth
//
//	@Param		shippingId	query		string	true	"query"
//
//	@Success	200			{object}	model.Shipping
//	@Router		/checkout/:shippingId [get]
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

// Checkout godoc
//
//	@Summary	purchase order
//	@Tags		shipping
//	@Produce	json
//	@Security	ApiKeyAuth
//
//	@Param		shipping	formData	model.Shipping	true	"formData"
//
//	@Success	200			{string}	string			"Checkout successfully"
//	@Router		/checkout [post]
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
