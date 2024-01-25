package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type PayOrderUsecase struct {
	OrderRepository repositories.OrderRepository
}

func (r *PayOrderUsecase) Execute(inputDto dtos.PayOrderDto) (*entities.Order, error) {
	orderToPay := r.OrderRepository.FindyId(inputDto.OrderId)

	if orderToPay.ID == 0 {
		return nil, errors.New("pedido não existe")
	}

	if orderToPay.Status != "waiting_payment" {
		return orderToPay, errors.New("pedido já está pago")
	}

	orderToPay.Status = "received"

	r.OrderRepository.Update(orderToPay)

	return orderToPay, nil
}
