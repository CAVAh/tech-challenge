package controllers

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
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
