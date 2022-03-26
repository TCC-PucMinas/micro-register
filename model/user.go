package model

import (
	"fmt"

	"github.com/TCC-PucMinas/micro-register/helpers"

	"github.com/TCC-PucMinas/micro-register/db"
)

type UsersPaginate struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
}

type User struct {
	Id         int64      `json:"id"`
	Phone      string     `json:"phone"`
	Forgot     string     `json:"forgot"`
	Business   string     `json:"business"`
	CpfCnpj    string     `json:"cpf_cnpj"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	CodeActive string     `json:"code_active"`
	Active     int        `json:"active"`
	Adress     []Address  `json:"addresses"`
	Permission Permission `json:"permission"`
}

func (u *User) GetOneUserByForgotCode() error {

	sql := db.ConnectDatabase()

	query := `select id from users where forgot = ? limit 1;`

	requestConfig, err := sql.Query(query, u.Forgot)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var id int64
		_ = requestConfig.Scan(&id)
		if id != 0 {
			u.Id = id
		}

	}

	return nil
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

func (u *User) GetOneUserByCodeActive() error {

	sql := db.ConnectDatabase()

	query := `select id from users where code_active = ? and active = ? limit 1;`

	user, err := sql.Query(query, u.CodeActive, u.Active)

	if err != nil {
		return err
	}

	for user.Next() {
		var id int64
		_ = user.Scan(&id)
		if id != 0 {
			u.Id = id
		}
	}

	return nil
}

func (u *User) UpdateUserSetActiveUser() error {
	sql := db.ConnectDatabase()

	query := "update users set code_active = ?, active = ? where code_active = ?"

	updateUser, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := updateUser.Exec(nil, u.Active, u.CodeActive)

	if e != nil {
		return e
	}

	return nil
}

func (u *User) UpdateUserSetNewPasswordByForgot() error {
	sql := db.ConnectDatabase()

	query := "update users set password = ?, forgot = ? where forgot = ?"

	updateUser, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := updateUser.Exec(u.Password, nil, u.Forgot)

	if e != nil {
		return e
	}

	return nil
}

func (u *User) CreateUser() (int64, error) {
	sql := db.ConnectDatabase()

	query := `insert into users (phone, business, cpf_cnpj, email, password, first_name, last_name, code_active, active)  values (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	insertUser, err := sql.Prepare(query)

	if err != nil {
		return 0, err
	}

	d, e := insertUser.Exec(u.Phone, u.Business, u.CpfCnpj, u.Email, u.Password, u.FirstName, u.LastName, u.CodeActive, u.Active)

	if e != nil {
		return 0, e
	}

	return d.LastInsertId()
}

func (u *User) GetOneUserByEmailAndCpf_Cpnj() error {

	sql := db.ConnectDatabase()

	query := `select id from users where email = ? or cpf_cnpj = ? limit 1;`

	user, err := sql.Query(query, u.Email, u.CpfCnpj)

	if err != nil {
		return err
	}

	for user.Next() {
		var id int64
		_ = user.Scan(&id)
		if id != 0 {
			u.Id = id
		}
	}

	return nil
}

func (u *User) GetByNameLike(name string, page, limit int64) (UsersPaginate, error) {
	var userArray UsersPaginate

	sql := db.ConnectDatabase()

	name = "%" + name + "%"
	paginate := helpers.Paginate{
		Page:  page,
		Limit: limit,
	}

	paginate.PaginateMounted()
	paginate.MountedQuery("users")

	query := fmt.Sprintf(`select u.id, u.first_name, u.last_name, u.email, u.phone, u.business, u.cpf_cnpj, p.name, %v 
										from users as u
											    join user_permissions as up on up.id_user = u.id
												join permissions as p on p.id = up.id
											where first_name like ? or last_name like ? LIMIT ? OFFSET ?;`, paginate.Query)

	requestConfig, err := sql.Query(query, name, name, paginate.Limit, paginate.Page)

	if err != nil {
		return userArray, err
	}

	for requestConfig.Next() {
		var userGet User
		var id int64
		var total int64
		var firstName, lastName, email, phone, business, cpfCnpj, permission string
		_ = requestConfig.Scan(&id, &firstName, &lastName, &email, &phone, &business, &cpfCnpj, &permission, &total)
		if id != 0 {
			userGet.Id = id
			userGet.FirstName = firstName
			userGet.LastName = lastName
			userGet.Email = email
			userGet.Phone = phone
			userGet.Business = business
			userGet.CpfCnpj = cpfCnpj
			userGet.Permission.Name = permission
			userArray.Total = total
			userArray.Users = append(userArray.Users, userGet)
		}
	}

	return userArray, nil
}
