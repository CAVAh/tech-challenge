package usecases

import (
	"errors"
	gateways2 "github.com/CAVAh/api-tech-challenge/src/adapters/gateways"
)

type DeleteProductUsecase struct {
	repository      gateways2.ProductRepository
	orderRepository gateways2.OrderRepository
}

func BuildDeleteProductUsecase(repository gateways2.ProductRepository, orderRepository gateways2.OrderRepository) *DeleteProductUsecase {
	return &DeleteProductUsecase{repository: repository, orderRepository: orderRepository}
}

func (p *DeleteProductUsecase) Execute(id uint) error {
	if p.orderRepository.ExistsOrderProduct(id) {
		return errors.New("product is associated with an order")
	}

	return p.repository.DeleteById(id)
}
