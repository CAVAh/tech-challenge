package dtos

type CreateCustomerDto struct {
	Name  string `json:"name"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
}
