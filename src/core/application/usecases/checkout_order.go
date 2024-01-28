package usecases

import (
	"errors"
	"time"

	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type CheckoutOrderUsecase struct {
	OrderRepository repositories.OrderRepository
}

func (r *CheckoutOrderUsecase) Execute(inputDto dtos.PayOrderDto) (*entities.Order, error) {
	orderToPay := r.OrderRepository.FindyId(inputDto.OrderId)

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
