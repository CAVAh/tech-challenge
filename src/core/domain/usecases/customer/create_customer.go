package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapter/gateways"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type CreateCustomerUsecase struct {
	CustomerRepository gateways.CustomerRepository
}

func (r *CreateCustomerUsecase) Execute(inputDto dtos.CreateCustomerDto) (*entities.Customer, error) {
	customer := entities.Customer{
		Name:  inputDto.Name,
		CPF:   inputDto.CPF,
		Email: inputDto.Email,
	}

	return r.CustomerRepository.Create(&customer)
}
