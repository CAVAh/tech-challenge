package controllers

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	order2 "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/order"
	usecases2 "github.com/CAVAh/api-tech-challenge/src/core/domain/usecases/payment"
	"github.com/CAVAh/api-tech-challenge/src/infra/db/repositories"
	mercado_pago2 "github.com/CAVAh/api-tech-challenge/src/infra/external/mercado_pago"
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
		PaymentInterface: &mercado_pago2.MercadoPagoIntegration{},
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
	var inputDto mercado_pago2.PostPayment

	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//value, _ := c.GetQuery("data.id")
	//orderId, _ := strconv.Atoi(value)
	//Explicação: para funcionar o teste do mercado livre, precisa pegar do ID,
	//já que o external reference não é mandado. Mas o id de dentro da aplicação estará em external reference
	var orderId, _ = strconv.Atoi(inputDto.AdditionalInfo.ExternalReference)

	usecase := order2.ChangeOrderStatusUsecase{
		OrderRepository: &repositories.OrderRepository{},
	}

	var statusChangeTo string
	if inputDto.State == mercado_pago2.Finished {
		statusChangeTo = enums.Received
	} else if inputDto.State == mercado_pago2.Error || inputDto.State == mercado_pago2.Canceled {
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
