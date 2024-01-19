package models

type OrderProduct struct {
	OrderId     uint `gorm:"primaryKey"`
	ProductId   uint `gorm:"primaryKey"`
	Quantity    int
	Observation string
}
