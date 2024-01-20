package controllers

import (
	"net/http"

	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/application/usecases"

	"github.com/gin-gonic/gin"
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
