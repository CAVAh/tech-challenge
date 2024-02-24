package controllers

import (
	dtos2 "github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	usecases2 "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases"
	"github.com/CAVAh/api-tech-challenge/src/db/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

func ListOrder(c *gin.Context) {
	status := c.Query("status")
	orderBy := c.Query("orderBy")
	sortBy := c.Query("sortBy")

	orderRepository := &repositories.OrderRepository{}

	usecase := usecases2.ListOrderUsecase{
		OrderRepository: orderRepository,
	}

	orders, err := usecase.Execute(sortBy, orderBy, status)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if orders == nil {
		c.JSON(http.StatusOK, []string{})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
	var inputDto dtos2.CreateOrderDto

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

	usecase := usecases2.CreateOrderUsecase{
		OrderRepository:    orderRepository,
		CustomerRepository: customerRepository,
		ProductRepository:  productRepository,
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

func CheckoutOrder(c *gin.Context) {
	var inputDto dtos2.PayOrderDto

	if err := c.BindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validator.Validate(inputDto); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	orderRepository := &repositories.OrderRepository{}

	usecase := usecases2.CheckoutOrderUsecase{
		OrderRepository: orderRepository,
	}

	order, err := usecase.Execute(inputDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    order.Status,
		"updatedAt": order.UpdatedAt,
	})
}
