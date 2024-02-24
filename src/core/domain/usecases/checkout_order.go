package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	"github.com/CAVAh/api-tech-challenge/src/gateways/repositories"
	"time"

	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type CheckoutOrderUsecase struct {
	OrderRepository repositories.OrderRepository
}

func (r *CheckoutOrderUsecase) Execute(inputDto dtos.PayOrderDto) (*entities.Order, error) {
	orderToPay := r.OrderRepository.FindById(inputDto.OrderId)

	if orderToPay.ID == 0 {
		return nil, errors.New("pedido não existe")
	}

	if orderToPay.Status != "created" {
		return orderToPay, errors.New("pedido já está confirmado")
	}

	orderToPay.Status = "received"
	orderToPay.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	r.OrderRepository.Update(orderToPay)

	return orderToPay, nil
}
