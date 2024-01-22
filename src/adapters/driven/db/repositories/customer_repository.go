package repositories

import (
	"errors"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/gorm"
	"github.com/CAVAh/api-tech-challenge/src/adapters/driven/db/models"
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"strings"
)

type CustomerRepository struct {
}

func (r CustomerRepository) Create(entity *entities.Customer) (*entities.Customer, error) {
	customer := models.Customer{
		Name:  entity.Name,
		CPF:   entity.CPF,
		Email: entity.Email,
	}

	if err := gorm.DB.Create(&customer).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("cliente já existe no sistema")
		} else {
			return nil, errors.New("ocorreu um erro desconhecido ao criar o cliente")
		}
	}

	result := customer.ToDomain()

	return &result, nil
}

func (r CustomerRepository) List(entity *entities.Customer) ([]entities.Customer, error) {
	var customers []models.Customer

	if cpf := entity.CPF; cpf != "" {
		gorm.DB.Where("cpf = ?", cpf).Find(&customers)
	} else {
		gorm.DB.Find(&customers)
	}

	var response []entities.Customer

	for _, customer := range customers {
		response = append(response, customer.ToDomain())
	}

	return response, nil
}

func (r CustomerRepository) FindFirstById(id int) (*entities.Customer, error) {
	var customers []models.Customer
	gorm.DB.Where("id = ?", id).Find(&customers)

	if len(customers) == 0 {
		return nil, nil
	} else {
		var entity = customers[0].ToDomain()
		return &(entity), nil
	}
}
