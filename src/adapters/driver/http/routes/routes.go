package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequests() {
	router := gin.Default()

	SetupCustomerRoutes(router)
	SetupProductRoutes(router)

	err := router.Run()

	if err != nil {
		log.Panic(err)
		return
	}
}
