package model

import (
	"errors"
	"log"
	"strconv"

	"github.com/TCC-PucMinas/micro-register/db"
)

var keyTokenRedisRoutes = "key-route"

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
				where r.path = ? limit 1;`

	requestConfig, err := sql.Query(query, id, r.Path)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var id, path, method string
		_ = requestConfig.Scan(&id, &path, &method)
		i64, _ := strconv.ParseInt(id, 10, 64)
		r.Id = i64
		r.Path = path
		r.Method = method
	}

	log.Println("r", r)
	if r.Id == 0 {
		return errors.New("Not found key")
	}

	return nil
}
