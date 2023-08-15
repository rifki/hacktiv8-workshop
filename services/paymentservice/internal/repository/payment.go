package repository

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rifki/hacktiv8-workshop/services/paymentservice/internal/model"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (repo *PaymentRepository) CreatePayment(ctx *gin.Context, paymentData model.Payment) (id int, err error) {
	db := ctx.MustGet("db").(*gorm.DB)
	// db := repo.db
	qx := db.Save(&paymentData)
	db.Save(&paymentData)
	if qx.Error != nil {
		log.Println("error:", qx.Error)
		err = fmt.Errorf(qx.Error.Error())
		return
	}

	return paymentData.ID, nil
}

func (repo *PaymentRepository) GetListPayment(ctx *gin.Context) (paymentList []model.Payment) {
	db := ctx.MustGet("db").(*gorm.DB)
	// db := repo.db
	db.Find(&paymentList)
	return paymentList
}
