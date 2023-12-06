package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name  string `json:"name"`
	CPF   string `json:"cpf" gorm:"unique;index"`
	Email string `json:"email" gorm:"unique;index"`
}
