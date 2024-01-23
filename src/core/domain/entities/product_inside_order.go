package entities

type ProductInsideOrder struct {
	Product     Product `json:"product"`
	Quantity    int     `json:"quantity"`
	Observation string  `json:"observation"`
}
