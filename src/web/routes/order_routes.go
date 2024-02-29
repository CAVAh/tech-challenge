package routes

import (
	orderController "github.com/CAVAh/api-tech-challenge/src/web/controllers/order"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("", orderController.CreateOrder)
		orderRoutes.GET("", orderController.ListOngoingOrders)
		orderRoutes.POST("/checkout", orderController.CheckoutOrder)
		orderRoutes.POST("/change-status", orderController.ChangeOrderStatus)
	}
}
