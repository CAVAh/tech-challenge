package routes

import (
	paymentController "github.com/CAVAh/api-tech-challenge/src/adapter/controllers/payment"
	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/payments")
	{
		orderRoutes.GET("/statuses", paymentController.CheckOrderPaymentStatus)
		orderRoutes.GET("/qr-code", paymentController.GetOrderQrCode)
		orderRoutes.POST("", paymentController.MercadoPagoPayment)
	}
}
