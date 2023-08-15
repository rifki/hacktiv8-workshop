package database

import (
	"fmt"
	"log"

	"github.com/rifki/hacktiv8-workshop/services/paymentservice/pkg/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

type IConnection interface {
	DbConnect() (db *gorm.DB, err error)
}

func New() IConnection {
	return &Connection{
		db: &gorm.DB{},
	}
}

func (conn *Connection) DbConnect() (db *gorm.DB, err error) {
	var (
		dbUser  = helper.GetEnv("PAYMENTSERVICE_DB_USER")
		dbPass  = helper.GetEnv("PAYMENTSERVICE_DB_PASSWORD")
		dbHost  = helper.GetEnv("PAYMENTSERVICE_DB_HOST")
		dbName  = helper.GetEnv("PAYMENTSERVICE_DB_NAME")
		dbPort  = helper.GetEnv("PAYMENTSERVICE_DB_PORT")
		TZ      = helper.GetEnv("PAYMENTSERVICE_TZ")
		sslMode = helper.GetEnv("PAYMENTSERVICE_SSL_MODE")
	)

	// dsn
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s`,
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
		sslMode,
		TZ,
	)

	// log.Print("dsn:", dsn)
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("Connected to database Failed:", err)
		return db, fmt.Errorf("Failed not connect DB: " + err.Error())
	}
	log.Println("Connected to database")

	return db, nil
}
