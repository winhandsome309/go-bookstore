// This file to handle request and response
package http

import (
	"go-bookstore/internal/product/model"
	"go-bookstore/internal/product/service"
	model_user "go-bookstore/internal/user/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

type ProductHandler struct {
	service service.IProductService
}

func NewProductHandler(service service.IProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// GetAllProduct godoc
//
//	@Summary	get all products
//	@Tags		products
//	@Produce	json
//	@Success	200	{array}		model.Product
//	@Success	200	{object}	model_user.User
//	@Router		/products [get]
func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	products, err := h.service.GetAllProduct(c)
	if err != nil {
		log.Error("Failed to get all product", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get all product"})
		return
	}
	// c.JSON(http.StatusOK, products)
	userJson, ok := c.Get("user")
	if ok {
		user := userJson.(model_user.User)
		c.JSON(http.StatusOK, gin.H{
			"products": products,
			"user": gin.H{
				"id":      user.Id,
				"email":   user.Email,
				"balance": user.Balance,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"products": products,
			"user":     nil,
		},
		)
	}
}

// GetProductById godoc
//
//	@Summary	get product by id
//	@Tags		products
//	@Produce	json
//
//	@Param		id	query	string	true	"query"
//
//	@Success	200	{array}	model.Product
//	@Router		/products/:id [get]
func (h *ProductHandler) GetProductById(c *gin.Context) {
	id := c.Param("id")
	productId, _ := strconv.Atoi(id)
	product, err := h.service.GetProductById(c, productId)
	if err != nil {
		log.Error("Product not found", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct godoc
//
//	@Summary	create new product
//	@Tags		products
//	@Produce	json
//
//	@Param		product	body		model.Product	true	"body"
//
//	@Success	200		{string}	string			"Create successfully"
//	@Success	200		{array}		model.Product
//	@Router		/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var productNew model.Product
	if err := c.ShouldBind(&productNew); c.Request.Body == nil || err != nil {
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

// UpdateProduct godoc
//
//	@Summary	update product info
//	@Tags		products
//	@Produce	json
//
//	@Param		id			query		string					true	"query"
//	@Param		requeset	formData	model.UpdateProductReq	true	"formData"
//
//	@Success	200			{string}	string					"Update successfully"
//	@Success	200			{array}		model.Product
//	@Router		/products/:id [patch]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req model.UpdateProductReq
	if err := c.ShouldBind(&req); c.Request.Body == nil || err != nil {
		log.Error("Failed to get body", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}
	productId, _ := strconv.Atoi(id)
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

// DeleteProduct godoc
//
//	@Summary	delete product
//	@Tags		products
//	@Produce	json
//
//	@Param		id	query		string	true	"query"
//
//	@Success	200	{string}	string	"Delete successfully"
//	@Router		/products/:id [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productId, _ := strconv.Atoi(id)
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
