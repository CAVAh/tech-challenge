package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ListOrderUsecase struct {
	OrderRepository repositories.OrderRepository
}

func (r *ListOrderUsecase) Execute(pageSize int64, pageNumber int64, status string) ([]entities.Order, error) {

	order, err := r.OrderRepository.List(pageSize, pageNumber, status)

	return order, err
}
