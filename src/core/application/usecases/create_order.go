package usecases

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/application/ports/repositories"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type CreateOrderUsecase struct {
	OrderRepository repositories.OrderRepository
}

func (r *CreateOrderUsecase) Execute(inputDto dtos.CreateOrderDto) (*entities.Order, error) {
	order := models.Order{
		CustomerId: inputDto.CustomerId,
	}

	return r.OrderRepository.Create(&order, inputDto.ProductIds)
}
