package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/TCC-PucMinas/micro-register/db"
)

var keyTokenRedisRouteByPathAndUserId = "key-route-path-and-by-user-id"

type Route struct {
	Id         int64      `json:"id"`
	Path       string     `json:"path"`
	Method     string     `json:"method"`
	Permission Permission `json:"permission"`
}

func setRedisCacheRouteByPathAndUserId(route *Route) error {
	db, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	json, err := json.Marshal(route)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyTokenRedisRouteByPathAndUserId, json)

	return db.Set(key, json, 1*time.Hour).Err()
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
		r = &route
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
		var id, path, method string
		_ = requestConfig.Scan(&id, &path, &method)
		i64, _ := strconv.ParseInt(id, 10, 64)
		r.Id = i64
		r.Path = path
		r.Method = method
	}

	if r.Id == 0 {
		return errors.New("Not found key")
	}

	_ = setRedisCacheRouteByPathAndUserId(r)
	return nil
}
