package routes

import (
	customerController "github.com/CAVAh/api-tech-challenge/src/adapters/driver/controllers/customer"
	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.Engine) {
	customerRoutes := router.Group("/customers")
	{
		customerRoutes.GET("", customerController.ListCustomers)
		customerRoutes.POST("", customerController.CreateCustomer)
		customerRoutes.GET("/:id", customerController.FindById)
		customerRoutes.DELETE("/:id", customerController.DeleteCustomer)
		customerRoutes.PATCH("/:id", customerController.UpdateCustomer)
	}
}
