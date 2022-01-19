package model

import (
	"errors"
	"strconv"

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
		var id, firstName, lastName, email, phone, bussines, password string
		_ = requestConfig.Scan(&id, &firstName, &lastName, &email, &phone, &bussines, &password)
		i64, _ := strconv.ParseInt(id, 10, 64)
		user.Id = i64
		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.Phone = phone
		user.Business = bussines
		user.Password = password
	}

	if user.Id == 0 {
		return user, errors.New("Not found key")
	}

	return user, nil
}
