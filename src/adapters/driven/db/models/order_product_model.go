package models

import (
	"gorm.io/gorm"
)

type OrderProduct struct {
	gorm.Model
	OrderID     uint // gorm know it is mapped to entity even without annotation
	ProductID   uint
	Quantity    int
	Observation string
	Product     Product `gorm:"foreignKey:ProductID;references:ID"`
}
