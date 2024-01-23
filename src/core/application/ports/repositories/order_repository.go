package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type OrderRepository interface {
	Create(order dtos.CreateOrderDto) (*entities.Order, error)
	List() ([]entities.Order, error)
}
