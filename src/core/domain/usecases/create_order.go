package usecases

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/enums"
	"github.com/CAVAh/api-tech-challenge/src/gateways"
)

type CreateOrderUsecase struct {
	OrderRepository    gateways.OrderRepository
	CustomerRepository gateways.CustomerRepository
	ProductRepository  gateways.ProductRepository
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
		Status:        enums.Created,
		PaymentStatus: enums.WaitingPayment,
		Customer:      entities.Customer{ID: inputDto.CustomerId},
		Products:      products,
	}

	return r.OrderRepository.Create(&entity)
}

func (r *CreateOrderUsecase) CustomerExists(id uint) error {
	customer, err := r.CustomerRepository.FindFirstById(id)

	if customer == nil {
		return errors.New("usuário não encontrado")
	} else if err != nil {
		return errors.New("algum erro desconhecido aconteceu ao procurar o usuário")
	} else {
		return nil
	}
}

func (r *CreateOrderUsecase) AllProductsExists(ids []uint) error {
	filteredIds := RemoveDuplicates(ids)
	products, err := r.ProductRepository.FindByIds(filteredIds)

	if len(products) != len(filteredIds) {
		return errors.New("algum dos produtos não foi encontrado")
	} else if err != nil {
		return errors.New("algum erro desconhecido aconteceu ao procurar os produtos")
	} else {
		return nil
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
	var errCustomer = r.CustomerExists(inputDto.CustomerId)
	if errCustomer != nil {
		return errCustomer
	}

	var errProducts = r.AllProductsExists(inputDto.GetProductIds())
	if errProducts != nil {
		return errProducts
	}

	return nil
}
