package routes

import (
	customerController "github.com/CAVAh/api-tech-challenge/src/adapters/driver/controllers/customer"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/customers", customerController.ListCustomers)
	r.POST("/customers", customerController.CreateCustomer)
	r.GET("/customers/:id", customerController.FindById)
	r.DELETE("/customers/:id", customerController.DeleteCustomer)
	r.PATCH("/customers/:id", customerController.UpdateCustomer)

	err := r.Run()

	if err != nil {
		log.Panic(err)
		return
	}
}
