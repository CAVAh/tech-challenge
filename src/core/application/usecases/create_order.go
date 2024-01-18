package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type CreateOrderUsecase struct {
	OrderRepository    repositories.OrderRepository
	CustomerRepository repositories.CustomerRepository
	ProductRepository  repositories.ProductRepository
}

func (r *CreateOrderUsecase) Execute(inputDto dtos.CreateOrderDto) (*entities.Order, error) {
	return r.OrderRepository.Create(inputDto.CustomerId, inputDto.ProductIds)
}

func (r *CreateOrderUsecase) CustomerExists(id int) bool {
	customer, err := r.CustomerRepository.FindFirstById(id)

	if err != nil || customer == nil {
		return false
	} else {
		return true
	}
}

func (r *CreateOrderUsecase) AllProductsExists(ids []int) bool {
	products, err := r.ProductRepository.FindById(ids)

	if err != nil || len(products) != len(ids) {
		return false
	} else {
		return true
	}
}
