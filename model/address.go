package model

import "github.com/TCC-PucMinas/micro-register/db"

type Address struct {
	Id         int64  `json:"id"`
	Street     string `json:"street"`
	State      string `json:"state"`
	Number     string `json:"number"`
	Country    string `json:"country"`
	Complement string `json:"complement"`
}

func (a *Address) CreateAddress() (int64, error) {
	sql := db.ConnectDatabase()

	query := `insert into addresses (street, state, number, country, complement) values (?, ?, ?, ?, ?);`

	insertUser, err := sql.Prepare(query)

	if err != nil {
		return 0, err
	}

	d, e := insertUser.Exec(a.Street, a.State, a.Number, a.Country, a.Complement)

	if e != nil {
		return 0, e
	}

	return d.LastInsertId()
}
