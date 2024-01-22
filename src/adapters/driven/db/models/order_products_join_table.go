package models

type OrderProduct struct {
	ID          uint `gorm:"primaryKey"`
	OrderID     uint `gorm:"primaryKey"`
	ProductID   uint `gorm:"primaryKey"`
	Quantity    int
	Observation string
}
