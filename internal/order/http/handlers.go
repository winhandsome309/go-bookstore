package http

import (
	"go-bookstore/internal/order/model"
	"go-bookstore/internal/order/service"
	model_user "go-bookstore/internal/user/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type OrderHandler struct {
	service service.IOrderService
}

func NewOrderHandler(service service.IOrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// GetOrder godoc
//
//	@Summary	get order of user
//	@Tags		orders
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{object}	model.OrderResponse
//	@Router		/orders [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	// Get user through cookie
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please sign in again",
		})
		return
	}
	customerId := user.(model_user.User).Id
	orderResponse, err := h.service.GetOrderResponse(c, customerId)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "get fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order": orderResponse,
	})
}

// DeleteOrder godoc
//
//	@Summary	delete order
//	@Tags		orders
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		orderId	query		string	true	"Query"
//
//	@Success	200		{string}	string	"delete	successfully"
//
//	@Router		/orders [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("orderId")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "delete fail",
		})
		return
	}
	err = h.service.DeleteOrder(c, orderId)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "delete fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete successfully",
	})
}

// GetOrderLines godoc
//
//	@Summary	get all orderlines of user
//	@Tags		orderlines
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{array}	model.OrderLine
//
//	@Router		/orderlines [get]
func (h *OrderHandler) GetOrderLines(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please sign in again",
		})
	}
	customerId := user.(model_user.User).Id
	orderLines, err := h.service.GetOrderLines(c, customerId)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"orderlines": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orderlines": orderLines,
	})
}

// CreateUpdateOrderLine godoc
//
//	@Summary	update or create orderline
//	@Tags		orders
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		orderLine	formData	model.OrderLine	true	"formData"
//	@Success	200			{string}	string			"add successfully"
//	@Router		/orders [post]
func (h *OrderHandler) CreateUpdateOrderLine(c *gin.Context) {
	var orderLine model.OrderLine
	if err := c.ShouldBind(&orderLine); c.Request.Body == nil || err != nil {
		log.Error("Failed to get body", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters",
		})
		return
	}
	user, ok := c.Get("user")
	customerId := user.(model_user.User).Id
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please sign in again",
		})
		return
	}
	err := h.service.CreateUpdateOrderLine(c, customerId, &orderLine)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "add fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "add successfully",
	})
}

// DeleteOrderLineById godoc
//
//	@Summary	delete orderlines by product_id and order_id
//	@Tags		orderlines
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		product_id	formData	string	true	"formData"
//	@Param		order_id	formData	string	true	"formData"
//	@Success	200			{string}	string	"remove successfully"
//
//	@Success	200			{object}	model.Order
//
//	@Router		/orderlines [delete]
func (h *OrderHandler) DeleteOrderLineById(c *gin.Context) {
	productId := c.PostForm("product_id")
	orderId := c.PostForm("order_id")
	productId_v, err := strconv.Atoi(productId)
	if err != nil {
		panic(err)
	}
	orderId_v, err := strconv.Atoi(orderId)
	if err != nil {
		panic(err)
	}
	order, err := h.service.DeleteOrderLineById(c, productId_v, orderId_v)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "remove fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "remove successfully",
		"order":   order,
	})
}
