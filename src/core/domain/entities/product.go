package entities

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryId  int     `json:"categoryId"`
	CreatedAt   string  `json:"createdAt"`
}
