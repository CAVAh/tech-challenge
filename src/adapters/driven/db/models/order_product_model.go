package models

import (
	"gorm.io/gorm"
)

type OrderProduct struct {
	gorm.Model
	OrderID     uint
	ProductID   uint
	Quantity    int
	Observation string
	Product     Product
}
