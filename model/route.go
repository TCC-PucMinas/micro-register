package model

import (
	"github.com/TCC-PucMinas/micro-register/db"
)

type Route struct {
	Id         int64      `json:"id"`
	Path       string     `json:"path"`
	Method     string     `json:"method"`
	Permission Permission `json:"permission"`
}

func (r *Route) GetOneRouteByPathAndUserId(id string) error {

	sql := db.ConnectDatabase()

	query := `select r.id, r.path, r.method from routes as r
					join permissions as p on r.id_permission = p.id
					join user_permissions as u on u.id_user = ?
				where r.path = ? and r.method = ? limit 1;`

	requestConfig, err := sql.Query(query, id, r.Path, r.Method)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var path, method string
		var id int64
		_ = requestConfig.Scan(&id, &path, &method)
		if id != 0 {
			r.Id = id
			r.Path = path
			r.Method = method
		}
	}
	return nil
}
