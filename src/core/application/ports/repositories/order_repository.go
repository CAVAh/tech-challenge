package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type OrderRepository interface {
	Create(order *models.Order, productIds []int) (*entities.Order, error)
	List() ([]entities.Order, error)
}
