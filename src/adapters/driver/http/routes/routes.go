package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	SetupCustomerRoutes(router)
	SetupProductRoutes(router)
	SetupOrderRoutes(router)
	SetupPaymentRoutes(router)

	err := router.Run()

	if err != nil {
		log.Panic(err)
		return
	}
}
