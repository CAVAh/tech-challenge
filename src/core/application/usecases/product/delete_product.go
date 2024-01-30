package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
)

type DeleteProductUsecase struct {
	repository      repositories.ProductRepository
	orderRepository repositories.OrderRepository
}

func BuildDeleteProductUsecase(repository repositories.ProductRepository, orderRepository repositories.OrderRepository) *DeleteProductUsecase {
	return &DeleteProductUsecase{repository: repository, orderRepository: orderRepository}
}

func (p *DeleteProductUsecase) Execute(id uint) error {
	if p.orderRepository.ExistsOrderProduct(id) {
		return errors.New("product is associated with an order")
	}

	return p.repository.DeleteById(id)
}
