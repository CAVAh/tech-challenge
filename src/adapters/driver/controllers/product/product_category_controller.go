package product

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/repositories"
	usecasesProduct "github.com/CAVAh/api-tech-challenge/src/core/application/usecases/product"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ListCategory(ctx *gin.Context) {
	productRepository := repositories.ProductCategoryRepository{}

	result, err := usecasesProduct.BuildListCategory(productRepository).Execute()

	if err != nil {
		log.Println("there was an error to retrieve products", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
