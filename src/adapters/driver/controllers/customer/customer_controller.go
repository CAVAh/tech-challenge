package controllers

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/application/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCustomers(c *gin.Context) {
	var inputDto dtos.ListCustomerDto

	if err := c.BindQuery(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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

// FindById TODO: Remover
func FindById(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	gorm.DB.First(&customer, id)

	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Cliente não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, customer)
}

// DeleteCustomer TODO: Remover
func DeleteCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	gorm.DB.Delete(&customer, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Cliente deletado com sucesso",
	})
}

// UpdateCustomer TODO: Remover
func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	gorm.DB.First(&customer, id)

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	gorm.DB.Model(&customer).UpdateColumns(customer)
	c.JSON(http.StatusOK, customer)
}
