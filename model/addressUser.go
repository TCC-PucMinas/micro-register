package model

import "github.com/TCC-PucMinas/micro-register/db"

type AddressUser struct {
	Id        int64 `json:"id"`
	UserId    int64 `json:"user_id"`
	AddressId int64 `json:"address_id"`
}

func (a *AddressUser) CreateAddressUser() (int64, error) {
	sql := db.ConnectDatabase()

	query := `insert into user_addresses (id_user, id_address) values (?, ?)`

	insertAddress, err := sql.Prepare(query)

	if err != nil {
		return 0, err
	}

	d, e := insertAddress.Exec(a.UserId, a.AddressId)

	if e != nil {
		return 0, e
	}

	return d.LastInsertId()
}
