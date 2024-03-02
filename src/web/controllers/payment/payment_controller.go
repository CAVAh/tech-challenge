package controllers

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	order2 "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/order"
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

func MercadoPagoPayment(c *gin.Context) {
	value, _ := c.GetQuery("data.id")
	orderId, _ := strconv.Atoi(value)

	var inputDto mercado_pago.PostPayment

	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	usecase := order2.ChangeOrderStatusUsecase{
		OrderRepository: &repositories.OrderRepository{},
	}

	var statusChangeTo string
	if inputDto.State == mercado_pago.Finished {
		statusChangeTo = enums.Received
	} else if inputDto.State == mercado_pago.Error || inputDto.State == mercado_pago.Canceled {
		statusChangeTo = enums.Cancelled
	}

	if statusChangeTo != "" {
		_, err := usecase.Execute(uint(orderId), statusChangeTo)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.Status(http.StatusOK)
}
