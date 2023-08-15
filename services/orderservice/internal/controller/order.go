package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifki/hacktiv8-workshop/services/orderservice/internal/dto"
	"github.com/rifki/hacktiv8-workshop/services/orderservice/internal/service"
)

func CreateOrder(ctx *gin.Context) {
	var reqOrder dto.RequestCreateOrder
	err := ctx.BindJSON(&reqOrder)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := service.NewOrderService().CreateOrder(ctx, reqOrder)
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

func GetOrder(ctx *gin.Context) {
	response := service.NewOrderService().GetListOrder(ctx)
	fmt.Println(response)

	ctx.JSON(http.StatusOK, response)
}
