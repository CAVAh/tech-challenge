package routes

import (
	customerController "github.com/CAVAh/api-tech-challenge/src/web/controllers/customer"
	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.Engine) {
	customerRoutes := router.Group("/customers")
	{
		customerRoutes.GET("", customerController.ListCustomers)
		customerRoutes.POST("", customerController.CreateCustomer)
	}
}
