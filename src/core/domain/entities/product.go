package entities

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryID  int     `json:"categoryId"`
	CreatedAt   string  `json:"createdAt"`
}

func NewProduct(id int, name string,
	price float64, description string, categoryID int, createdAt string) *Product {
	return &Product{Id: id,
		Name:        name,
		Price:       price,
		Description: description,
		CategoryID:  categoryID,
		CreatedAt:   createdAt}
}

func (p Product) IsExistingProduct() bool {
	return p.Id > 0
}

func (p *Product) PatchFields(name string, price float64, description string, categoryId int) {
	p.Name = name
	p.Price = price
	p.Description = description
	p.CategoryID = categoryId
}
