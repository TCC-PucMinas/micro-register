package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/TCC-PucMinas/micro-register/db"
)

var (
	keyTokenRedisRoleByUserIdId = "key-role-by-id"
)

type Permission struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func setRedisCacheRolesByUserId(permission []Permission) error {
	db := db.ConnectDatabaseRedis()

	json, err := json.Marshal(permission)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyTokenRedisRoleByUserIdId, json)

	return db.Set(key, json, 1*time.Hour).Err()
}

func getRedisCacheRolesByUserId(userId int64) ([]Permission, error) {
	userPermissons := []Permission{}

	redis := db.ConnectDatabaseRedis()

	key := fmt.Sprintf("%v - %v", keyTokenRedisRoleByUserIdId, userId)

	value, err := redis.Get(key).Result()

	if err != nil {
		return userPermissons, err
	}

	if err := json.Unmarshal([]byte(value), &userPermissons); err != nil {
		return userPermissons, err
	}

	return userPermissons, nil
}

func GetAllRoles(id int64) ([]Permission, error) {

	if permissions, err := getRedisCacheRolesByUserId(id); err == nil {
		return permissions, err
	}

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

	_ = setRedisCacheRolesByUserId(userPermissons)

	return userPermissons, nil
}
