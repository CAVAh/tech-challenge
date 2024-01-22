package routes

import (
	controllersProduct "github.com/CAVAh/api-tech-challenge/src/adapters/driver/controllers/product"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine) {
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("", controllersProduct.List)
		productRoutes.POST("", controllersProduct.Create)
		productRoutes.PATCH("/:id", controllersProduct.Update)
		productRoutes.DELETE("/:id", controllersProduct.Delete)
		productRoutes.GET("/:id", controllersProduct.Read)
		productRoutes.GET("/categories", controllersProduct.ListCategory)
	}
}
