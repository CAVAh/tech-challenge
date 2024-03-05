package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
)

type CheckPaymentStatusUsecase struct {
	OrderRepository gateways.OrderRepository
}

func (r *CheckPaymentStatusUsecase) Execute(orderId uint) (string, error) {
	order := r.OrderRepository.FindById(orderId)

	if order.ID == 0 {
		return "Pedido não encontado", errors.New("pedido não encontrado")
	}
	//TODO: fazer FindByiD retornar erro, essa comparação não deveria ser feita

	if order.PaymentStatus == enums.Paid {
		return "Pedido pago", nil
	} else if order.PaymentStatus == enums.AwaitingPayment {
		return "Pedido aguardando pagamento", nil
	} else {
		return "Status desconhecido", nil
	}
}
