package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
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
		CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
