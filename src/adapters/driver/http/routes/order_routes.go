package routes

import (
	orderController "github.com/CAVAh/api-tech-challenge/src/adapters/driver/controllers/order"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("", orderController.CreateOrder)
		orderRoutes.GET("", orderController.ListOrder)
	}
}
