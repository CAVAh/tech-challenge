package usecases

import (
	"errors"
	"time"

	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type PayOrderUsecase struct {
	OrderRepository repositories.OrderRepository
}

func (r *PayOrderUsecase) Execute(orderId int64) (entities.Order, error) {

	orderToPay := r.OrderRepository.FindyId(orderId)

	if orderToPay.Status != "waiting_payment" {
		return orderToPay, errors.New("Pedido já pago")
	}

	orderToPay.Status = "received"
	orderToPay.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	r.OrderRepository.Update(orderToPay)

	return orderToPay, nil
}
