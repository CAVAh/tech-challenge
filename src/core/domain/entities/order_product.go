package entities

type OrderProduct struct {
	ProductID   uint   `json:"id" validate:"nonzero"`
	Quantity    int    `json:"quantity" validate:"nonzero"`
	Observation string `json:"observation"`
}
