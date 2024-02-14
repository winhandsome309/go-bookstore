package repository

import (
	"go-bookstore/internal/order/model"
	model_product "go-bookstore/internal/product/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	// Order
	CreateOrder(c *gin.Context, order *model.Order) error
	GetOrderByCusId(c *gin.Context, customerId int) (*model.Order, error)
	GetOrderByOrderId(c *gin.Context, orderId int) (*model.Order, error)
	UpdateOrder(c *gin.Context, order *model.Order) error
	// Orderline
	CreateOrderLine(c *gin.Context, orderLine *model.OrderLine) error
	GetOrderLine(c *gin.Context, productId int, orderId int) (*model.OrderLine, error)
	GetOrderLines(c *gin.Context, orderId int) (*[]model.OrderLine, error)
	GetOrderLineById(c *gin.Context, orderLineId int) (*model.OrderLine, error)
	UpdateOrderLine(c *gin.Context, orderLine *model.OrderLine) error
	DeleteOrderLine(c *gin.Context, orderLine *model.OrderLine) error
	// Product
	GetProductById(c *gin.Context, id int) (*model_product.Product, error)
}

type OrderReposity struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderReposity {
	return &OrderReposity{db: db}
}

func (r *OrderReposity) CreateOrder(c *gin.Context, order *model.Order) error {
	err := r.db.Create(order).Error
	return err
}

func (r *OrderReposity) GetOrderByCusId(c *gin.Context, customerId int) (*model.Order, error) {
	var order model.Order
	err := r.db.Where("customer_id = ?", customerId).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderReposity) GetOrderByOrderId(c *gin.Context, orderId int) (*model.Order, error) {
	var order model.Order
	err := r.db.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		return &model.Order{}, err
	}
	return &order, nil
}

func (r *OrderReposity) UpdateOrder(c *gin.Context, order *model.Order) error {
	err := r.db.Save(order).Error
	return err
}

func (r *OrderReposity) CreateOrderLine(c *gin.Context, orderLine *model.OrderLine) error {
	var orderLineId int
	r.db.Raw("SELECT NEXTVAL(pg_get_serial_sequence('order_lines', 'id'))").Scan(&orderLineId)
	err := r.db.Create(orderLine).Error
	return err
}

func (r *OrderReposity) GetOrderLine(c *gin.Context, productId int, orderId int) (*model.OrderLine, error) {
	var orderLine model.OrderLine
	err := r.db.Where("product_id = ? AND order_id = ?", productId, orderId).First(&orderLine).Error
	return &orderLine, err
}

func (r *OrderReposity) GetOrderLines(c *gin.Context, orderId int) (*[]model.OrderLine, error) {
	var orderLines []model.OrderLine
	err := r.db.Where("order_id = ?", orderId).Find(&orderLines).Error
	if err != nil {
		return nil, err
	}
	return &orderLines, nil
}

func (r *OrderReposity) GetOrderLineById(c *gin.Context, orderLineId int) (*model.OrderLine, error) {
	var orderLine model.OrderLine
	err := r.db.Where("id = ?", orderLineId).First(&orderLine).Error
	if err != nil {
		return nil, err
	}
	return &orderLine, nil
}

func (r *OrderReposity) UpdateOrderLine(c *gin.Context, orderLine *model.OrderLine) error {
	err := r.db.Save(orderLine).Error
	return err
}

func (r *OrderReposity) DeleteOrderLine(c *gin.Context, orderLine *model.OrderLine) error {
	err := r.db.Delete(&orderLine).Error
	return err
}

func (r *OrderReposity) GetProductById(c *gin.Context, id int) (*model_product.Product, error) {
	var product model_product.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
