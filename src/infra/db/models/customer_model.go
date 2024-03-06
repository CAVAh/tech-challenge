package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"github.com/CAVAh/api-tech-challenge/src/utils"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name  string
	CPF   string `gorm:"unique;index"`
	Email string `gorm:"unique;index"`
}

func (c Customer) ToDomain() entities.Customer {
	return entities.Customer{
		ID:        c.ID,
		Name:      c.Name,
		CPF:       c.CPF,
		Email:     c.Email,
		CreatedAt: c.CreatedAt.Format(utils.CompleteEnglishDateFormat),
	}
}
