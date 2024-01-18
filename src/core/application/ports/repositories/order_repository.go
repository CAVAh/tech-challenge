package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type OrderRepository interface {
	Create(costumerId int, productIds []int) (*entities.Order, error)
	List() ([]entities.Order, error)
}
