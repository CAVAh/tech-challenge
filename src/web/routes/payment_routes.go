package routes

import (
	paymentController "github.com/CAVAh/api-tech-challenge/src/web/controllers/payment"
	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/payments")
	{
		orderRoutes.GET("/check-order-status", paymentController.CheckOrderPaymentStatus)
		orderRoutes.GET("/order-qr-code", paymentController.GetOrderQrCode)
		orderRoutes.POST("/mercado-pago-payment", paymentController.MercadoPagoPayment)
	}
}
