package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/gateways"
)

type DeleteProductUsecase struct {
	repository      gateways.ProductRepository
	orderRepository gateways.OrderRepository
}

func BuildDeleteProductUsecase(repository gateways.ProductRepository, orderRepository gateways.OrderRepository) *DeleteProductUsecase {
	return &DeleteProductUsecase{repository: repository, orderRepository: orderRepository}
}

func (p *DeleteProductUsecase) Execute(id uint) error {
	if p.orderRepository.ExistsOrderProduct(id) {
		return errors.New("product is associated with an order")
	}

	return p.repository.DeleteById(id)
}
