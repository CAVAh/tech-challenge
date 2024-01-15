package entities

type Order struct {
	ID        uint      `json:"id"`
	CreatedAt string    `json:"createdAt"`
	Customer  Customer  `json:"customer"`
	Products  []Product `json:"products"`
}
