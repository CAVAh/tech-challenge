package controllers

import (
	dtos2 "github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	usecases2 "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases"
	"github.com/CAVAh/api-tech-challenge/src/db/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

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

func CheckoutOrder(c *gin.Context) { //TODO: can be deleted, is the same as ChangeOrderStatus with received param
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

func CheckOrderPaymentStatus(c *gin.Context) {
	value, _ := c.GetQuery("orderId")
	orderId, _ := strconv.Atoi(value)

	orderRepository := &repositories.OrderRepository{}

	usecase := usecases2.CheckPaymentStatusUsecase{
		OrderRepository: orderRepository,
	}

	response, err := usecase.Execute(uint(orderId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func ChangeOrderStatus(c *gin.Context) {
	var inputDto dtos2.ChangeOrderStatusDto

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

	usecase := usecases2.ChangeOrderStatusUsecase{
		OrderRepository: orderRepository,
	}

	order, err := usecase.Execute(inputDto.OrderId, inputDto.ChangeToStatus)

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

func ListOngoingOrders(c *gin.Context) {
	orderRepository := &repositories.OrderRepository{}

	usecase := usecases2.ListOngoingOrdersUsecase{
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
