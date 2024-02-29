package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/gateways"
)

type ListCustomerUsecase struct {
	CustomerRepository gateways.CustomerRepository
}

func (r *ListCustomerUsecase) Execute(inputDto dtos.ListCustomerDto) ([]entities.Customer, error) {
	customer := entities.Customer{
		CPF: inputDto.CPF,
	}

	return r.CustomerRepository.List(&customer)
}
