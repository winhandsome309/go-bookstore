package http

import (
	"go-bookstore/internal/product/service"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	products, err := h.service.GetAllProduct(c)
	if err != nil {
		log.Error("Failed to get all product", err)
		return
	}
	c.JSON(http.StatusOK, products)
}
