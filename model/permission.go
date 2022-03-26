package model

import (
	"github.com/TCC-PucMinas/micro-register/db"
)

type Permission struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetAllRoles(id int64) ([]Permission, error) {

	sql := db.ConnectDatabase()

	userPermissons := []Permission{}

	p := Permission{}

	query := `select p.id as id, p.name as name from permissions as p
				inner join user_permissions as u on p.id = u.id_permission
		 	  where u.id_user = ?`

	requestConfig, err := sql.Query(query, id)

	if err != nil {
		return userPermissons, err
	}

	for requestConfig.Next() {
		var name string
		var id int64
		_ = requestConfig.Scan(&id, &name)
		if id != 0 {
			p.Id = id
			p.Name = name
			userPermissons = append(userPermissons, p)
		}
	}

	return userPermissons, nil
}
