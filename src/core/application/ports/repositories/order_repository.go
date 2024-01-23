package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/core/application/dtos"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type OrderRepository interface {
	List(sortBy string, orderBy string, status string) ([]entities.Order, error)
	FindyId(orderId uint) *entities.Order
	Update(*entities.Order)
	Create(order dtos.CreateOrderDto) (*entities.Order, error)
}
