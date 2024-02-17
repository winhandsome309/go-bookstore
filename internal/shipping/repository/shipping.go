package repository

import (
	"fmt"
	model_order "go-bookstore/internal/order/model"
	model_product "go-bookstore/internal/product/model"
	"go-bookstore/internal/shipping/model"
	model_user "go-bookstore/internal/user/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IShippingRepository interface {
	GetAllShipping(c *gin.Context) (*[]model.Shipping, error)
	GetShippingById(c *gin.Context, shippingId string) (*model.Shipping, error)
	Checkout(c *gin.Context, shipping *model.Shipping) error
	GetOrderById(c *gin.Context, orderId string) (*model_order.Order, error)
	GetUserById(c *gin.Context, userId string) (*model_user.User, error)
	UpdateUser(c *gin.Context, user *model_user.User) error
	DeleteOrder(c *gin.Context, order *model_order.Order) error
	GetOrderLineById(c *gin.Context, orderLineId string) (*model_order.OrderLine, error)
	DeleteOrderLineById(c *gin.Context, orderLine *model_order.OrderLine) error
	GetProductById(c *gin.Context, id int) (*model_product.Product, error)
	UpdateProduct(c *gin.Context, product *model_product.Product) error
	DeleteProduct(c *gin.Context, product *model_product.Product) error
}

type ShippingRepository struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) *ShippingRepository {
	return &ShippingRepository{db: db}
}

func (r *ShippingRepository) GetAllShipping(c *gin.Context) (*[]model.Shipping, error) {
	var shippings []model.Shipping
	err := r.db.Find(&shippings).Error
	return &shippings, err
}

func (r *ShippingRepository) GetShippingById(c *gin.Context, shippingId string) (*model.Shipping, error) {
	var shipping model.Shipping
	err := r.db.Where("id = ?", shippingId).First(&shipping).Error
	return &shipping, err
}

func (r *ShippingRepository) Checkout(c *gin.Context, shipping *model.Shipping) error {
	err := r.db.Create(shipping).Error
	fmt.Println(err)
	return err
}

func (r *ShippingRepository) GetOrderById(c *gin.Context, orderId string) (*model_order.Order, error) {
	var order model_order.Order
	err := r.db.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *ShippingRepository) GetUserById(c *gin.Context, userId string) (*model_user.User, error) {
	var user model_user.User
	err := r.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *ShippingRepository) UpdateUser(c *gin.Context, user *model_user.User) error {
	err := r.db.Save(&user).Error
	return err
}

func (r *ShippingRepository) DeleteOrder(c *gin.Context, order *model_order.Order) error {
	err := r.db.Delete(&order).Error
	return err
}

func (r *ShippingRepository) GetOrderLineById(c *gin.Context, orderLineId string) (*model_order.OrderLine, error) {
	var orderLine model_order.OrderLine
	err := r.db.Where("id = ?", orderLineId).First(&orderLine).Error
	return &orderLine, err
}

func (r *ShippingRepository) DeleteOrderLineById(c *gin.Context, orderLine *model_order.OrderLine) error {
	err := r.db.Delete(orderLine).Error
	return err
}

func (r *ShippingRepository) GetProductById(c *gin.Context, id int) (*model_product.Product, error) {
	var product model_product.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ShippingRepository) UpdateProduct(c *gin.Context, product *model_product.Product) error {
	err := r.db.Save(&product).Error
	return err
}

func (r *ShippingRepository) DeleteProduct(c *gin.Context, product *model_product.Product) error {
	err := r.db.Delete(&product).Error
	return err
}
