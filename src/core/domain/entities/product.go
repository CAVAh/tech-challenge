package entities

type Product struct {
	id          uint    `json:"id"`
	name        string  `json:"name"`
	price       float64 `json:"price"`
	description string  `json:"description"`
	categoryID  uint    `json:"categoryId"`
	createdAt   string  `json:"createdAt"`
}

func NewProduct(id uint, name string,
	price float64, description string, categoryID uint, createdAt string) *Product {
	return &Product{id: id,
		name:        name,
		price:       price,
		description: description,
		categoryID:  categoryID,
		createdAt:   createdAt}
}

func (p *Product) ID() uint {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) CategoryID() uint {
	return p.categoryID
}

func (p *Product) CreatedAt() string {
	return p.createdAt
}

func (p Product) IsExistingProduct() bool {
	return p.id > 0
}

func (p *Product) PatchFields(name string, price float64, description string, categoryId uint) {
	p.name = name
	p.price = price
	p.description = description
	p.categoryID = categoryId
}
