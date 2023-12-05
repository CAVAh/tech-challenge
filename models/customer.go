package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name  string `json:"nome"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
}
