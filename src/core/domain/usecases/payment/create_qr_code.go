package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/gateways"
)

type CreateQrCodeUsecase struct {
	PaymentInterface gateways.PaymentInterface
	OrderRepository  gateways.OrderRepository
}

func (r *CreateQrCodeUsecase) Execute(orderId uint) (string, error) {
	order := r.OrderRepository.FindById(orderId)
	return r.PaymentInterface.CreatePayment(*order)
}
