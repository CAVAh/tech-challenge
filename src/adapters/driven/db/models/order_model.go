package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID uint
	Customer   Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Products   []OrderProduct
	Status     string
}

//PS: instead of using toDomain, use OrderModelToOrderEntity
