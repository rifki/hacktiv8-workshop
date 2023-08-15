package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rifki/hacktiv8-workshop/services/orderservice/internal/controller"
	"github.com/rifki/hacktiv8-workshop/services/orderservice/pkg/database"
	"github.com/rifki/hacktiv8-workshop/services/orderservice/pkg/helper"
)

///
func init() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("start")
	db, err := database.New().DbConnect()
	if err != nil {
		panic("database not connect")
	}

	port := helper.GetEnv("PORT_ORDER")

	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(database.InjectDB(db))
	app.UseRawPath = true
	app.UnescapePathValues = true
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"OPTIONS", "PUT", "POST", "GET", "DELETE"}
	app.Use(cors.New(config))

	app.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := app.Group("/v1")
	{
		v1.POST("/order", controller.CreateOrder)
		v1.GET("/order", controller.GetOrder)
	}
	app.Run(":" + port)
}
