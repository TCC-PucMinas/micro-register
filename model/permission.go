package model

import (
	"strconv"

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
		var id, name string
		_ = requestConfig.Scan(&id, &name)
		i64, _ := strconv.ParseInt(id, 10, 64)
		p.Id = i64
		p.Name = name

		userPermissons = append(userPermissons, p)
	}

	return userPermissons, nil
}
