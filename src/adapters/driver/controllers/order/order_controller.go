package controllers

import (
	"net/http"
	"strconv"

	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/application/usecases"

	"github.com/gin-gonic/gin"
)

func ListOrder(c *gin.Context) {
	pageSizeString := c.Query("pageSize")
	pageNumberString := c.Query("pageNumber")
	status := c.Query("status")

	pageSize, err := strconv.ParseInt(pageSizeString, 10, 64)
	pageNumber, err := strconv.ParseInt(pageNumberString, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"erro": "Parâmetro 'numero' inválido"})
		return
	}

	orderRepository := &repositories.OrderRepository{}

	usecase := usecases.ListOrderUsecase{
		OrderRepository: orderRepository,
	}

	orders, err := usecase.Execute(pageSize, pageNumber, status)

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
