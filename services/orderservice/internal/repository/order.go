package repository

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rifki/hacktiv8-workshop/services/orderservice/internal/dto"
	"github.com/rifki/hacktiv8-workshop/services/orderservice/internal/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (repo *OrderRepository) CreateOrder(ctx *gin.Context, orderData model.Order) (id int, err error) {
	db := ctx.MustGet("db").(*gorm.DB)
	// db := repo.db
	qx := db.Save(&orderData)
	if qx.Error != nil {
		log.Println("error:", qx.Error)
		err = fmt.Errorf(qx.Error.Error())
		return
	}

	return orderData.ID, nil
}

func (repo *OrderRepository) GetListOrder(ctx *gin.Context) (orderList []dto.Order) {
	db := ctx.MustGet("db").(*gorm.DB)
	// db := repo.db
	db.Model(model.Order{}).Find(&orderList)
	return orderList
}
