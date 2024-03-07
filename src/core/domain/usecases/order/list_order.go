package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ListOrderUsecase struct {
	OrderRepository gateways.OrderRepository
}

func (r *ListOrderUsecase) Execute(sortBy string, orderBy string, status string) ([]entities.Order, error) {
	order, err := r.OrderRepository.List(sortBy, orderBy, status)

	return order, err
}
