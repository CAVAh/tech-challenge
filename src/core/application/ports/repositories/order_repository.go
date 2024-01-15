package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type OrderRepository interface {
	List(pageSize int64, pageNumber int64, status string) ([]entities.Order, error)
}
