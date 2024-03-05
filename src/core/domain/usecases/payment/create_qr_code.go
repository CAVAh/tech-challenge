package usecases

import (
	gateways2 "github.com/CAVAh/api-tech-challenge/src/adapter/gateways"
)

type CreateQrCodeUsecase struct {
	PaymentInterface gateways2.PaymentInterface
	OrderRepository  gateways2.OrderRepository
}

func (r *CreateQrCodeUsecase) Execute(orderId uint) (string, error) {
	order := r.OrderRepository.FindById(orderId)
	return r.PaymentInterface.CreatePayment(*order)
}
