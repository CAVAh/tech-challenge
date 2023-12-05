package controllers

import (
	"net/http"

	"github.com/CAVAh/api-tech-challenge/database"
	"github.com/CAVAh/api-tech-challenge/models"
	"github.com/gin-gonic/gin"
)

func ListCustomers(c *gin.Context) {
	var list []models.Customer
	database.DB.Find(&list)
	c.JSON(200, list)
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Create(&customer)
	c.JSON(http.StatusOK, customer)
}

func FindById(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	database.DB.First(&customer, id)

	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Cliente não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	database.DB.Delete(&customer, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Cliente deletado com sucesso",
	})
}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	database.DB.First(&customer, id)

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&customer).UpdateColumns(customer)
	c.JSON(http.StatusOK, customer)
}

func FindCustomerByCpf(c *gin.Context) {
	var customer models.Customer
	cpf := c.Param("cpf")
	database.DB.Where(&models.Customer{CPF: cpf}).First(&customer)

	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Customer não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, customer)
}
