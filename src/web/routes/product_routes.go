package routes

import (
	"github.com/CAVAh/api-tech-challenge/src/web/controllers/product"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine) {
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("", controllers.List)
		productRoutes.POST("", controllers.Create)
		productRoutes.PATCH("/:id", controllers.Update)
		productRoutes.DELETE("/:id", controllers.Delete)
		productRoutes.GET("/:id", controllers.Read)
		productRoutes.GET("/categories", controllers.ListCategory)
	}
}
