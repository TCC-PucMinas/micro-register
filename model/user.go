package model

import "github.com/TCC-PucMinas/micro-register/db"

type User struct {
	Id        int64     `json:"id"`
	Phone     string    `json:"phone"`
	Forgot    string    `json:"forgot"`
	Business  string    `json:"business"`
	CpfCnpj   string    `json:"cpf_cnpj"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Adress    []Address `json:"addresses"`
}

func (u *User) UpdateUserForgotById() (bool, error) {
	sql := db.ConnectDatabase()

	query := "update users set forgot = ? where id = ?"

	updateHealthyCheck, err := sql.Prepare(query)

	if err != nil {
		return false, err
	}

	_, e := updateHealthyCheck.Exec(u.Forgot, u.Id)

	if e != nil {
		return false, e
	}

	return true, nil
}
