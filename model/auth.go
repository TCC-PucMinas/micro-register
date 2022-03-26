package model

import (
	"github.com/TCC-PucMinas/micro-register/db"
)

type Authenticate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (auth *Authenticate) GetOneUserByEmail() (User, error) {

	sql := db.ConnectDatabase()

	user := User{}

	query := `select id, first_name, last_name, email, phone, business, password from users where email = ? limit 1;`

	requestConfig, err := sql.Query(query, auth.Email)

	if err != nil {
		return user, err
	}

	for requestConfig.Next() {
		var firstName, lastName, email, phone, bussines, password string
		var id int64
		_ = requestConfig.Scan(&id, &firstName, &lastName, &email, &phone, &bussines, &password)
		if id != 0 {
			user.Id = id
			user.FirstName = firstName
			user.LastName = lastName
			user.Email = email
			user.Phone = phone
			user.Business = bussines
			user.Password = password
		}
	}

	return user, nil
}
