package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"github.com/CAVAh/api-tech-challenge/src/utils"
	"time"

	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type CheckoutOrderUsecase struct {
	OrderRepository gateways.OrderRepository
}

func (r *CheckoutOrderUsecase) Execute(inputDto dtos.PayOrderDto) (*entities.Order, error) {
	orderToPay := r.OrderRepository.FindById(inputDto.OrderId)

	if orderToPay.ID == 0 {
		return nil, errors.New("pedido não existe")
	}
	//TODO: fazer FindByiD retornar erro, essa comparação não deveria ser feita

	if orderToPay.Status != enums.Created {
		return orderToPay, errors.New("pedido já está confirmado")
	}

	orderToPay.Status = enums.Received
	orderToPay.PaymentStatus = enums.Paid
	orderToPay.UpdatedAt = time.Now().Format(utils.CompleteEnglishDateFormat)

	r.OrderRepository.Update(orderToPay)

	return orderToPay, nil
}
