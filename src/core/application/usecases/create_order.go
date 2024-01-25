package usecases

import (
	"errors"
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

	var err = r.Verifications(inputDto)
	if err != nil {
		return nil, err
	}

	for _, p := range inputDto.Products {
		var productEntity = entities.Product{ID: p.Id}
		var productInsideOrder = entities.ProductInsideOrder{Quantity: p.Quantity, Observation: p.Observation, Product: productEntity}
		products = append(products, productInsideOrder)
	}

	var entity = entities.Order{
		Status:   "waiting_payment",
		Customer: entities.Customer{ID: inputDto.CustomerId},
		Products: products,
	}

	return r.OrderRepository.Create(&entity)
}

func (r *CreateOrderUsecase) CustomerExists(id uint) error {
	customer, err := r.CustomerRepository.FindFirstById(id)

	if err != nil || customer == nil {
		return errors.New("erro ao encontrar o usuário")
	} else {
		return err
	}
}

func (r *CreateOrderUsecase) AllProductsExists(ids []uint) error {
	filteredIds := RemoveDuplicates(ids)
	products, err := r.ProductRepository.FindByIds(filteredIds)

	if len(products) != len(filteredIds) {
		return errors.New("erro ao encontrar o produto")
	} else {
		return err
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

func (r *CreateOrderUsecase) Verifications(inputDto dtos.CreateOrderDto) error {
	if r.CustomerExists(inputDto.CustomerId) != nil {
		return errors.New("o cliente informado não existe!")
	}

	if r.AllProductsExists(inputDto.GetProductIds()) != nil {
		return errors.New("algum dos produtos não foram encontrados!")
	}

	return nil
}
