// This file to handle request and response
package http

import (
	"go-bookstore/internal/product/service"
	"net/http"

	"go-bookstore/internal/product/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ProductHandler struct {
	service service.IProductService
}

func NewProductHandler(service service.IProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	products, err := h.service.GetAllProduct(c)
	if err != nil {
		log.Error("Failed to get all product", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get all product"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductById(c *gin.Context) {
	productId := c.Param("id")
	product, err := h.service.GetProductById(c, productId)
	if err != nil {
		log.Error("Product ("+productId+") not found", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product (" + productId + ") not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var productNew model.Product
	if err := c.ShouldBindJSON(&productNew); c.Request.Body == nil || err != nil {
		log.Error("Failed to get body", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameters",
		})
	}
	err := h.service.CreateProduct(c, &productNew)
	if err != nil {
		log.Error("Create failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Create successfully",
		"product": productNew,
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productId := c.Param("id")
	var req model.UpdateProductReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		log.Error("Failed to get body", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}
	product, err := h.service.UpdateProduct(c, productId, &req)
	if err != nil {
		log.Error("Update failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update successfully",
		"product": product,
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productId := c.Param("id")
	err := h.service.DeleteProduct(c, productId)
	if err != nil {
		log.Error("Delete failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete successfully",
	})
}
