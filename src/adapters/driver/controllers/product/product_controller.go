package product

import (
	"fmt"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/repositories"
	dtosProduct "github.com/CAVAh/api-tech-challenge/src/core/application/dtos/product"
	usecasesProduct "github.com/CAVAh/api-tech-challenge/src/core/application/usecases/product"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	value, _ := ctx.GetQuery("categoryId")
	categoryId, _ := strconv.ParseUint(value, 0, 64)

	fmt.Println("\n\n\n\n\n\n param ", categoryId)

	productRepository := repositories.ProductRepository{}

	result, err := usecasesProduct.BuildList(productRepository).Execute(uint(categoryId))

	if err != nil {
		log.Println("there was an error to retrieve products", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func Create(c *gin.Context) {
	var inputDto dtosProduct.PersistProductDto

	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validator.Validate(inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.(validator.ErrorMap),
		})
		return
	}

	productRepository := repositories.ProductRepository{}

	usecase := usecasesProduct.BuildCreate(productRepository)

	result, err := usecase.Execute(inputDto)

	if err != nil {
		log.Println("there was an error to save the product", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Read(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Params.ByName("id"), 0, 64)

	useCase := usecasesProduct.BuildRead(repositories.ProductRepository{})

	product, err := useCase.Execute(uint(id))

	if !product.IsExistingProduct() {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	if err != nil {
		log.Println("there was an error to retrieve a product", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func Update(ctx *gin.Context) {
	var inputDto dtosProduct.PersistProductDto
	id, _ := strconv.ParseUint(ctx.Params.ByName("id"), 0, 64)

	if err := ctx.ShouldBindJSON(&inputDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	useCase := usecasesProduct.BuildEdit(repositories.ProductRepository{})

	product, err := useCase.Execute(inputDto, uint(id))

	if err != nil {
		log.Println("there was an error to retrieve a product", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if !product.IsExistingProduct() {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Params.ByName("id"), 0, 64)

	useCase := usecasesProduct.BuildDelete(repositories.ProductRepository{})

	err := useCase.Execute(uint(id))

	if err != nil {
		log.Println("there was an error to retrieve a product", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
