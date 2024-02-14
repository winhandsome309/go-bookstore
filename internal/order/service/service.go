package service

import (
	"go-bookstore/internal/order/model"
	"go-bookstore/internal/order/repository"
	model_product "go-bookstore/internal/product/model"

	"go-bookstore/pkg/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type IOrderService interface {
	GetOrder(c *gin.Context, customerId int) (*model.Order, error)
	GetOrderResponse(c *gin.Context, customerId int) (*model.OrderResponse, error)
	DeleteOrder(c *gin.Context, orderId int) error
	GetOrderByOrderId(c *gin.Context, orderId int) (*model.Order, error)
	GetOrderLines(c *gin.Context, customerId int) (*[]model.OrderLine, error)
	CreateUpdateOrderLine(c *gin.Context, customerId int, orderLine *model.OrderLine) error
	DeleteOrderLineById(c *gin.Context, productId int, orderId int) (*model.Order, error)
	GetProductById(c *gin.Context, productId int) (*model_product.Product, error)
}

type OrderService struct {
	repo repository.IOrderRepository
}

func NewOrderService(repo repository.IOrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetOrder(c *gin.Context, customerId int) (*model.Order, error) {
	order, err := s.repo.GetOrderByCusId(c, customerId)
	// Order not exist => Create new order
	if err != nil {
		order = &model.Order{CustomerID: customerId, TotalItem: 0, TotalPrice: 0, Lines: nil}
		err = s.repo.CreateOrder(c, order)
		if err != nil {
			log.Error(err)
			return nil, nil
		}
	}
	return order, nil
}

func (s *OrderService) GetOrderResponse(c *gin.Context, customerId int) (*model.OrderResponse, error) {
	order, err := s.repo.GetOrderByCusId(c, customerId)
	var orderResponse model.OrderResponse
	// Order not exist => Create new order
	if err != nil {
		order = &model.Order{CustomerID: customerId, TotalItem: 0, TotalPrice: 0, Lines: nil}
		err = s.repo.CreateOrder(c, order)
		if err != nil {
			log.Error(err)
			return nil, nil
		}
		utils.Merge(&orderResponse, &order)
	} else {
		utils.Merge(&orderResponse, &order)
		// Delete first element (null)
		orderResponse.Lines = nil
		for _, val := range order.Lines {
			orderLine, err := s.repo.GetOrderLineById(c, val)
			if err != nil {
				return nil, err
			}
			orderResponse.Lines = append(orderResponse.Lines, orderLine)
		}
	}
	return &orderResponse, nil
}

func (s *OrderService) DeleteOrder(c *gin.Context, orderId int) error {
	order, err := s.repo.GetOrderByOrderId(c, orderId)
	if err != nil {
		return err
	}
	for _, val := range order.Lines {
		orderLine, err := s.repo.GetOrderLineById(c, val)
		if err != nil {
			return err
		}
		s.repo.DeleteOrderLine(c, orderLine)
	}
	order.Lines = nil
	order.TotalItem = 0
	order.TotalPrice = 0
	err = s.repo.UpdateOrder(c, order)
	return err
}

func (s *OrderService) GetOrderByOrderId(c *gin.Context, orderId int) (*model.Order, error) {
	order, err := s.repo.GetOrderByOrderId(c, orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) GetOrderLines(c *gin.Context, customerId int) (*[]model.OrderLine, error) {
	order, err := s.GetOrder(c, customerId)
	if err != nil {
		return nil, err
	}
	orderLines, err := s.repo.GetOrderLines(c, order.Id)
	if err != nil {
		return nil, err
	}
	return orderLines, nil
}

func (s *OrderService) CreateUpdateOrderLine(c *gin.Context, customerId int, orderLine *model.OrderLine) error {
	order, err := s.GetOrder(c, customerId)
	if err != nil {
		return err
	}
	orderLine.OrderId = order.Id
	orderLineOld, err := s.repo.GetOrderLine(c, orderLine.ProductId, order.Id)
	// Update order
	order.TotalItem += orderLine.Quantity
	order.TotalPrice += orderLine.Price
	// If orderline not exist => Create new one
	if err != nil {
		err = s.repo.CreateOrderLine(c, orderLine)
		if err != nil {
			return err
		}
		order.Lines = append(order.Lines, orderLine.Id)
		s.repo.UpdateOrder(c, order)
		return nil
	}
	s.repo.UpdateOrder(c, order)
	// if orderline existed => Update orderline
	orderLineOld.Quantity += orderLine.Quantity
	orderLineOld.Price += orderLine.Price
	return s.repo.UpdateOrderLine(c, orderLineOld)
}

func (s *OrderService) DeleteOrderLineById(c *gin.Context, productId int, orderId int) (*model.Order, error) {
	order, err := s.GetOrderByOrderId(c, orderId)
	if err != nil {
		return nil, err
	}
	orderLine, err := s.repo.GetOrderLine(c, productId, orderId)
	if err != nil {
		return nil, err
	}
	// Update order
	order.TotalItem -= orderLine.Quantity
	order.TotalPrice -= orderLine.Price
	for i := range order.Lines {
		if order.Lines[i] == orderLine.Id {
			order.Lines[i] = order.Lines[len(order.Lines)-1]
			order.Lines = order.Lines[:len(order.Lines)-1]
			break
		}
	}
	err = s.repo.UpdateOrder(c, order)
	if err != nil {
		return nil, err
	}
	err = s.repo.DeleteOrderLine(c, orderLine)
	return order, err
}

func (s *OrderService) GetProductById(c *gin.Context, productId int) (*model_product.Product, error) {
	product, err := s.repo.GetProductById(c, productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}
