package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rifki/hacktiv8-workshop/services/orderservice/internal/dto"
	"github.com/rifki/hacktiv8-workshop/services/orderservice/internal/model"
	"github.com/rifki/hacktiv8-workshop/services/orderservice/internal/repository"
)

type OrderService struct {
	repository repository.OrderRepository
}

type IOrderService interface {
	CreateOrder(ctx *gin.Context, reqOrder dto.RequestCreateOrder) (id int, err error)
	GetListOrder(ctx *gin.Context) []dto.Order
}

func NewOrderService() IOrderService {
	return &OrderService{}
}

func (service *OrderService) CreateOrder(ctx *gin.Context, reqOrder dto.RequestCreateOrder) (id int, err error) {
	orderModel := model.Order{
		Price:         reqOrder.Price,
		CustomerID:    reqOrder.CustomerID,
		Items:         reqOrder.Items,
		PaymentStatus: "UNPAID",
	}
	fmt.Println(orderModel)
	id, err = service.repository.CreateOrder(ctx, orderModel)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (service *OrderService) GetListOrder(ctx *gin.Context) []dto.Order {
	return service.repository.GetListOrder(ctx)
}
