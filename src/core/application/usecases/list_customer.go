package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type ListCustomerUsecase struct {
	CustomerRepository repositories.CustomerRepository
}

func (r *ListCustomerUsecase) Execute(inputDto dtos.ListCustomerDto) ([]entities.Customer, error) {
	customer := entities.Customer{
		CPF: inputDto.CPF,
	}

	return r.CustomerRepository.List(&customer)
}
