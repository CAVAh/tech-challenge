package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type CustomerRepository interface {
	Create(customer *entities.Customer) (*entities.Customer, error)
	List(customer *entities.Customer) ([]entities.Customer, error)
}
