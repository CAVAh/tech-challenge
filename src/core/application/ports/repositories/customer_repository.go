package repositories

import (
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
)

type CustomerRepository interface {
	Create(customer *models.Customer) (*entities.Customer, error)
	List(customer *models.Customer) ([]entities.Customer, error)
}
