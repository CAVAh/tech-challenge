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
	return r.OrderRepository.Create(inputDto)
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
	filteredIds := RemoveDuplicates(ids)

	products, err := r.ProductRepository.FindById(filteredIds)

	if err != nil || len(products) != len(filteredIds) {
		return false
	} else {
		return true
	}
}

func RemoveDuplicates(ids []int) []int { // TODO: move to utils
	bucket := make(map[int]bool)
	var result []int
	for _, id := range ids {
		if _, ok := bucket[id]; !ok {
			bucket[id] = true
			result = append(result, id)
		}
	}

	return result
}
