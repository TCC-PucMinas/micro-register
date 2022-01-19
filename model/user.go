package model

type User struct {
	Id        int64     `json:"id"`
	Phone     string    `json:"phone"`
	Business  string    `json:"business"`
	CpfCnpj   string    `json:"cpf_cnpj"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Adress    []Address `json:"addresses"`
}
