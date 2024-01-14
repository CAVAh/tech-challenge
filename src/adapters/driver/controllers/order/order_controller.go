package controllers

import (
	"net/http"

	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/application/usecases"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

func CreateOrder(c *gin.Context) {
	var inputDto dtos.CreateOrderDto

	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validator.Validate(inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	//TODO: fazer as validacoes necessarias

	//customerRepository := &repositories.CustomerRepository{}

	//usecase := usecases.CreateCustomerUsecase{
	//	CustomerRepository: customerRepository,
	//}

	//result, err := usecase.Execute(inputDto)

	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}

	c.JSON(http.StatusOK, inputDto)
}

func CreateProduct(c *gin.Context) {
	var inputDto dtos.CreateProductDto

	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validator.Validate(inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	productRepository := &repositories.ProductRepository{}

	usecase := usecases.CreateProductUsecase{
		ProductRepository: productRepository,
	}

	result, err := usecase.Execute(inputDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ListProducts(c *gin.Context) {
	productRepository := &repositories.ProductRepository{}

	products, err := productRepository.List()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
