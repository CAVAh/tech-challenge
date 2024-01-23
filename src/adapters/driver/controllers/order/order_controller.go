package controllers

import (
  "github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/application/usecases"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
)

func ListOrder(c *gin.Context) {
	status := c.Query("status")

	orderBy := c.Query("orderBy")
	sortBy := c.Query("sortBy")

	orderRepository := &repositories.OrderRepository{}

	usecase := usecases.ListOrderUsecase{
		OrderRepository: orderRepository,
	}

	orders, err := usecase.Execute(sortBy, orderBy, status)
  
  if orders == nil {
		c.JSON(http.StatusOK, []string{})
  }
}

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

	orderRepository := &repositories.OrderRepository{}
	customerRepository := &repositories.CustomerRepository{}
	productRepository := &repositories.ProductRepository{}

	usecase := usecases.CreateOrderUsecase{
		OrderRepository:    orderRepository,
		CustomerRepository: customerRepository,
		ProductRepository:  productRepository,
	}

	if !usecase.CustomerExists(inputDto.CustomerId) {
		c.JSON(http.StatusBadRequest, "Usuário não existe, não foi possível criar pedido")
		return
	}

	if !usecase.AllProductsExists(inputDto.GetProductIds()) {
		c.JSON(http.StatusBadRequest, "Algum dos produtos não existe, não foi possível criar pedido")
		return
	}

	result, err := usecase.Execute(inputDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func ListOrder(c *gin.Context) {
	orderRepository := &repositories.OrderRepository{}

	orders, err := orderRepository.List()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}
