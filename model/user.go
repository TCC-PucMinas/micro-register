package model

import (
	"errors"
	"strconv"

	"github.com/TCC-PucMinas/micro-register/db"
)

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

func (u *User) GetOneUserByForgotCode() error {

	sql := db.ConnectDatabase()

	query := `select id  from users where forgot = ? limit 1;`

	requestConfig, err := sql.Query(query, u.Forgot)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var id string
		_ = requestConfig.Scan(&id)
		i64, _ := strconv.ParseInt(id, 10, 64)
		u.Id = i64
	}

	if u.Id == 0 {
		return errors.New("Not found key")
	}

	return nil
}

func (u *User) UpdateUserSetNewPasswordByForgot() error {
	sql := db.ConnectDatabase()

	query := "update users set password = ?, forgot = ? where forgot = ?"

	updateHealthyCheck, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := updateHealthyCheck.Exec(u.Password, nil, u.Forgot)

	if e != nil {
		return e
	}

	return nil
}
