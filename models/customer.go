package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name  string
	CPF   string `gorm:"unique;index"`
	Email string `gorm:"unique;index"`
}
