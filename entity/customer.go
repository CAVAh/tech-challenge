package entity

type Customer struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}
