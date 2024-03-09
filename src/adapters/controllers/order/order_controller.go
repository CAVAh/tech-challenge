package controllers

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	order "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/order"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
	"strconv"
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

	orderRepository := &repositories.OrderRepository{}
	customerRepository := &repositories.CustomerRepository{}
	productRepository := &repositories.ProductRepository{}

	usecase := order.CreateOrderUsecase{
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

func ChangeOrderStatus(c *gin.Context) {
	var inputDto dtos.ChangeOrderStatusDto
	id, _ := strconv.Atoi(c.Params.ByName("id"))

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

	usecase := order.ChangeOrderStatusUsecase{
		OrderRepository: orderRepository,
	}

	orderResult, err := usecase.Execute(uint(id), inputDto.ChangeToStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orderId":   orderResult.ID,
		"status":    orderResult.Status,
		"updatedAt": orderResult.UpdatedAt,
	})
}

func ListOngoingOrders(c *gin.Context) {
	orderRepository := &repositories.OrderRepository{}

	usecase := order.ListOngoingOrdersUsecase{
		OrderRepository: orderRepository,
	}

	orders, err := usecase.Execute()

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
