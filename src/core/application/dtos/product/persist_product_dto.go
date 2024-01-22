package product

type PersistProductDto struct {
	ID          int
	Name        string  `json:"name" validate:"nonzero"`
	Price       float64 `json:"price" validate:"nonzero"`
	Description string  `json:"description" validate:"nonzero"`
	CategoryID  int     `json:"categoryId" validate:"nonzero"`
}
