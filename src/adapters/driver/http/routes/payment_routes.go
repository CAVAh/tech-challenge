package routes

import (
	controllers "github.com/CAVAh/api-tech-challenge/src/adapters/driver/controllers/payment"
	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/payments")
	{
		orderRoutes.POST("", controllers.PayOrder)
	}
}
