package service

import (
	"go-bookstore/internal/shipping/model"
	"go-bookstore/internal/shipping/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IShippingService interface {
	GetAllShipping(c *gin.Context) (*[]model.Shipping, error)
	GetShippingById(c *gin.Context, shippingId string) (*model.Shipping, error)
	Checkout(c *gin.Context, shipping *model.Shipping) error
}

type ShippingService struct {
	repo repository.IShippingRepository
}

func NewShippingService(repo repository.IShippingRepository) *ShippingService {
	return &ShippingService{repo: repo}
}

func (s *ShippingService) GetAllShipping(c *gin.Context) (*[]model.Shipping, error) {
	shippings, err := s.repo.GetAllShipping(c)
	if err != nil {
		return nil, err
	}
	return shippings, nil
}

func (s *ShippingService) GetShippingById(c *gin.Context, shippingId string) (*model.Shipping, error) {
	shipping, err := s.repo.GetShippingById(c, shippingId)
	if err != nil {
		return nil, err
	}
	return shipping, nil
}

func (s *ShippingService) Checkout(c *gin.Context, shipping *model.Shipping) error {
	err := s.repo.Checkout(c, shipping)
	order, err := s.repo.GetOrderById(c, strconv.Itoa(shipping.OrderId))
	if err != nil {
		return err
	}
	user, err := s.repo.GetUserById(c, strconv.Itoa(order.CustomerID))
	if err != nil {
		return err
	}
	user.Balance -= order.TotalPrice
	err = s.repo.UpdateUser(c, user)
	for _, orderLineId := range order.Lines {
		orderLine, err := s.repo.GetOrderLineById(c, strconv.Itoa(orderLineId))
		if err != nil {
			return err
		}
		err = s.repo.DeleteOrderLineById(c, orderLine)
		if err != nil {
			return err
		}
	}
	err = s.repo.DeleteOrder(c, order)
	return err
}
