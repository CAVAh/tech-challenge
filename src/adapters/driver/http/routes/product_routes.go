package routes

import (
	productController "github.com/CAVAh/api-tech-challenge/src/adapters/driver/controllers/product"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/product")
	{
		orderRoutes.POST("", productController.CreateProduct)
		orderRoutes.GET("", productController.ListProducts)
	}
}
