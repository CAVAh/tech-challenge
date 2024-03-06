package controllers

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	usecases "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/customer"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
)

func ListCustomers(c *gin.Context) {
	var inputDto dtos.ListCustomerDto

	if err := c.BindQuery(&inputDto); err != nil {
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

	customerRepository := &repositories.CustomerRepository{}

	usecase := usecases.ListCustomerUsecase{
		CustomerRepository: customerRepository,
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

func CreateCustomer(c *gin.Context) {
	var inputDto dtos.CreateCustomerDto

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

	customerRepository := &repositories.CustomerRepository{}

	usecase := usecases.CreateCustomerUsecase{
		CustomerRepository: customerRepository,
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
