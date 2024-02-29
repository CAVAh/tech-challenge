package gateways

import "github.com/CAVAh/api-tech-challenge/src/core/domain/entities"

type PaymentInterface interface {
	CreatePayment(order entities.Order) (string, error)
}
