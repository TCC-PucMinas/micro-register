package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TCC-PucMinas/micro-register/db"
)

var (
	keyTokenRedisRouteByPathAndUserId = "key-route-path-and-by-user-id"
	keyTimeRoute                      = time.Hour * 1
)

type Route struct {
	Id         int64      `json:"id"`
	Path       string     `json:"path"`
	Method     string     `json:"method"`
	Permission Permission `json:"permission"`
}

func setRedisCacheRouteByPathAndUserId(userId string, route *Route) error {
	db, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	json, err := json.Marshal(route)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyTokenRedisRouteByPathAndUserId, userId)

	return db.Set(key, json, keyTimeRoute).Err()
}

func getRedisCacheRouteByPathAndUserId(userId string) (Route, error) {
	route := Route{}

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return route, err
	}

	key := fmt.Sprintf("%v - %v", keyTokenRedisRouteByPathAndUserId, userId)

	value, err := redis.Get(key).Result()

	if err != nil {
		return route, err
	}

	if err := json.Unmarshal([]byte(value), &route); err != nil {
		return route, err
	}

	return route, nil
}

func (r *Route) GetOneRouteByPathAndUserId(id string) error {

	if route, err := getRedisCacheRouteByPathAndUserId(id); err == nil {
		r.Id = route.Id
		r.Method = route.Method
		r.Path = route.Path
		return nil
	}

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

	if r.Id != 0 {
		_ = setRedisCacheRouteByPathAndUserId(id, r)
	}

	return nil
}
