package controllers

import (
	dtos2 "github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	order2 "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/order"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	repositories2 "github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"net/http"
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

	orderRepository := &repositories2.OrderRepository{}
	customerRepository := &repositories.CustomerRepository{}
	productRepository := &repositories2.ProductRepository{}

	usecase := order2.CreateOrderUsecase{
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

	orderRepository := &repositories2.OrderRepository{}

	usecase := order2.CheckoutOrderUsecase{
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

	orderRepository := &repositories2.OrderRepository{}

	usecase := order2.ChangeOrderStatusUsecase{
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
	orderRepository := &repositories2.OrderRepository{}

	usecase := order2.ListOngoingOrdersUsecase{
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
