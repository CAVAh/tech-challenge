package usecases

import (
	"errors"
	repositories2 "github.com/CAVAh/api-tech-challenge/src/gateways/repositories"
)

type DeleteProductUsecase struct {
	repository      repositories2.ProductRepository
	orderRepository repositories2.OrderRepository
}

func BuildDeleteProductUsecase(repository repositories2.ProductRepository, orderRepository repositories2.OrderRepository) *DeleteProductUsecase {
	return &DeleteProductUsecase{repository: repository, orderRepository: orderRepository}
}

func (p *DeleteProductUsecase) Execute(id uint) error {
	if p.orderRepository.ExistsOrderProduct(id) {
		return errors.New("product is associated with an order")
	}

	return p.repository.DeleteById(id)
}
