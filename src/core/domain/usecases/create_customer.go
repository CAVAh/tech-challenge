package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/gateways/repositories"
)

type CreateCustomerUsecase struct {
	CustomerRepository repositories.CustomerRepository
}

func (r *CreateCustomerUsecase) Execute(inputDto dtos.CreateCustomerDto) (*entities.Customer, error) {
	customer := entities.Customer{
		Name:  inputDto.Name,
		CPF:   inputDto.CPF,
		Email: inputDto.Email,
	}

	return r.CustomerRepository.Create(&customer)
}
