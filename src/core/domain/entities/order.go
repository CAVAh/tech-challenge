package entities

type Order struct {
	ID        uint      `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	Status    string    `json:"status"`
	Customer  Customer  `json:"customer"`
	Products  []Product `json:"products"`
}
