package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"github.com/CAVAh/api-tech-challenge/src/gateways/repositories"
	"slices"
)

type ListOngoingOrdersUsecase struct {
	OrderRepository repositories.OrderRepository
}

func (r *ListOngoingOrdersUsecase) Execute() ([]entities.Order, error) {
	var createdAt = r.OrderRepository.GetCreatedAtFieldName()
	var ascOrder = r.OrderRepository.GetAscOrder()

	ready, err := r.OrderRepository.List(createdAt, ascOrder, enums.Ready)
	preparation, err := r.OrderRepository.List(createdAt, ascOrder, enums.Preparation)
	received, err := r.OrderRepository.List(createdAt, ascOrder, enums.Received)

	result := slices.Concat(ready, preparation, received)

	return result, err
}
