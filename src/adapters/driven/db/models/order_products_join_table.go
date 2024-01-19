package models

type OrderProduct struct {
	ID          int `gorm:"primaryKey"`
	OrderId     int `gorm:"primaryKey"`
	ProductId   int `gorm:"primaryKey"`
	Quantity    int
	Observation string
}
