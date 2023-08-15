package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rifki/hacktiv8-workshop/services/paymentservice/internal/dto"
	"github.com/rifki/hacktiv8-workshop/services/paymentservice/internal/model"
	"github.com/rifki/hacktiv8-workshop/services/paymentservice/internal/repository"
)

type PaymentService struct {
	repository repository.PaymentRepository
}

type IPaymentService interface {
	CreatePayment(ctx *gin.Context, reqOrder dto.RequestCreatePayment) (id int, err error)
	GetListPayment(ctx *gin.Context) []model.Payment
}

func NewPaymentService() IPaymentService {
	return &PaymentService{}
}

func (service *PaymentService) CreatePayment(ctx *gin.Context, reqPayment dto.RequestCreatePayment) (id int, err error) {
	paymentModel := model.Payment{
		OrderID:       reqPayment.OrderID,
		PaymentStatus: "PAID",
	}
	fmt.Println(paymentModel)

	// validate order

	id, err = service.repository.CreatePayment(ctx, paymentModel)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (service *PaymentService) GetListPayment(ctx *gin.Context) []model.Payment {
	return service.repository.GetListPayment(ctx)
}

func (service *PaymentService) Validate(ctx *gin.Context) {

}
