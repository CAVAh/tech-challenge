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
	var products []entities.ProductInsideOrder

	for _, p := range inputDto.Products {
		var productEntity = entities.Product{ID: p.Id}
		var productInsideOrder = entities.ProductInsideOrder{Quantity: p.Quantity, Observation: p.Observation, Product: productEntity}
		products = append(products, productInsideOrder)
	}

	var aa = entities.Order{
		Status:   "waiting_payment",
		Customer: entities.Customer{ID: inputDto.CustomerId},
		Products: products,
	}

	return r.OrderRepository.Create(&aa)
}

func (r *CreateOrderUsecase) CustomerExists(id uint) bool {
	customer, err := r.CustomerRepository.FindFirstById(id)

	if err != nil || customer == nil {
		return false
	} else {
		return true
	}
}

func (r *CreateOrderUsecase) AllProductsExists(ids []uint) bool {
	filteredIds := RemoveDuplicates(ids)

	products, err := r.ProductRepository.FindByIds(filteredIds)

	if err != nil || len(products) != len(filteredIds) {
		return false
	} else {
		return true
	}
}

func RemoveDuplicates(ids []uint) []uint { // TODO: move to utils
	bucket := make(map[uint]bool)
	var result []uint
	for _, id := range ids {
		if _, ok := bucket[id]; !ok {
			bucket[id] = true
			result = append(result, id)
		}
	}

	return result
}
