package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifki/hacktiv8-workshop/services/paymentservice/internal/dto"
	"github.com/rifki/hacktiv8-workshop/services/paymentservice/internal/service"
)

func CreatePayment(ctx *gin.Context) {
	var reqPayment dto.RequestCreatePayment
	err := ctx.BindJSON(&reqPayment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := service.NewPaymentService().CreatePayment(ctx, reqPayment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": result,
	})
}

func GetPayment(ctx *gin.Context) {
	response := service.NewPaymentService().GetListPayment(ctx)
	fmt.Println(response)

	ctx.JSON(http.StatusOK, response)
}
