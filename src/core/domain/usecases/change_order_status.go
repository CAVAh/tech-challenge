package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"github.com/CAVAh/api-tech-challenge/src/gateways/repositories"
	"github.com/CAVAh/api-tech-challenge/src/utils"
	"time"
)

type ChangeOrderStatusUsecase struct {
	OrderRepository repositories.OrderRepository
}

func (r *ChangeOrderStatusUsecase) Execute(orderId uint, changeToStatus string) (*entities.Order, error) {
	var err error
	order := r.OrderRepository.FindById(orderId)

	if order.ID == 0 {
		return nil, errors.New("pedido não existe")
	}
	//TODO: fazer FindByiD retornar erro, essa comparação não deveria ser feita

	switch changeToStatus {
	case enums.Cancelled:
		order, err = ChangeToCancelled(order)
	case enums.Received:
		order, err = ChangeToReceived(order)
	case enums.Preparation:
		order, err = ChangeToPreparation(order)
	case enums.Ready:
		order, err = ChangeToReady(order)
	case enums.Finished:
		order, err = ChangeToFinished(order)
	default:
		return order, errors.New("status desconhecido")
	}

	order.UpdatedAt = time.Now().Format(utils.CompleteEnglishDateFormat)
	r.OrderRepository.Update(order)

	return order, err
}

func ChangeToCancelled(order *entities.Order) (*entities.Order, error) {
	if order.Status != enums.Created || order.PaymentStatus != enums.WaitingPayment {
		return order, errors.New("não é possível cancelar o pedido")
	}

	order.Status = enums.Received
	order.PaymentStatus = enums.Paid

	return order, nil
}

func ChangeToReceived(order *entities.Order) (*entities.Order, error) {
	if order.Status != enums.Created || order.PaymentStatus != enums.WaitingPayment {
		return order, errors.New("não é possível mudar o pedido para Recebido")
	}

	order.Status = enums.Received
	order.PaymentStatus = enums.Paid

	return order, nil
}

func ChangeToPreparation(order *entities.Order) (*entities.Order, error) {
	if order.Status != enums.Received || order.PaymentStatus != enums.Paid {
		return order, errors.New("não é possível mudar o pedido para Em preparação")
	}

	order.Status = enums.Preparation

	return order, nil
}

func ChangeToReady(order *entities.Order) (*entities.Order, error) {
	if order.Status != enums.Preparation || order.PaymentStatus != enums.Paid {
		return order, errors.New("não é possível mudar o pedido para Pronto")
	}

	order.Status = enums.Ready

	return order, nil
}

func ChangeToFinished(order *entities.Order) (*entities.Order, error) {
	if order.Status != enums.Ready || order.PaymentStatus != enums.Paid {
		return order, errors.New("não é possível mudar o pedido para Finalizado")
	}

	order.Status = enums.Finished

	return order, nil
}
