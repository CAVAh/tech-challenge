package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type OrderRepository interface {
	List(sortBy string, orderBy string, status string) ([]entities.Order, error)
	FindById(orderId uint) *entities.Order
	Update(*entities.Order)
	Create(order *entities.Order) (*entities.Order, error)
	ExistsOrderProduct(productId uint) bool
}
