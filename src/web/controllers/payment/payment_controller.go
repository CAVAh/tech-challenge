package controllers

import (
	usecases2 "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/payment"
	"github.com/CAVAh/api-tech-challenge/src/db/repositories"
	"github.com/CAVAh/api-tech-challenge/src/external/mercado_pago"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckOrderPaymentStatus(c *gin.Context) {
	value, _ := c.GetQuery("orderId")
	orderId, _ := strconv.Atoi(value)

	usecase := usecases2.CheckPaymentStatusUsecase{
		OrderRepository: &repositories.OrderRepository{},
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

func GetOrderQrCode(c *gin.Context) {
	value, _ := c.GetQuery("orderId")
	orderId, _ := strconv.Atoi(value)

	usecase := usecases2.CreateQrCodeUsecase{
		PaymentInterface: &mercado_pago.MercadoPagoIntegration{},
		OrderRepository:  &repositories.OrderRepository{},
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
