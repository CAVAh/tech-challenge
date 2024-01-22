package controllers

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"net/http"

	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/application/usecases"
	"gopkg.in/validator.v2"

	"github.com/gin-gonic/gin"
)

func PayOrder(c *gin.Context) {
	var inputDto dtos.PayOrderDto

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

	usecase := usecases.PayOrderUsecase{
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
