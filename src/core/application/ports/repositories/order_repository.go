package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type OrderRepository interface {
	List(sortBy string, orderBy string, status string) ([]entities.Order, error)
	FindyId(orderId int64) *entities.Order
	Update(*entities.Order)
}
