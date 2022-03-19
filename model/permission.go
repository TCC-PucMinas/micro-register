package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TCC-PucMinas/micro-register/db"
)

var (
	keyTokenRedisRoleByUserIdId = "key-role-by-id"
	keyTimePermission           = time.Hour * 1
)

type Permission struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func setRedisCacheRolesByUserId(userId int64, permission []Permission) error {
	db, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	json, err := json.Marshal(permission)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyTokenRedisRoleByUserIdId, userId)

	return db.Set(key, json, keyTimePermission).Err()
}

func getRedisCacheRolesByUserId(userId int64) ([]Permission, error) {
	userPermissons := []Permission{}

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return userPermissons, err
	}

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
		var name string
		var id int64
		_ = requestConfig.Scan(&id, &name)
		if id != 0 {
			p.Id = id
			p.Name = name
			userPermissons = append(userPermissons, p)
		}
	}

	if len(userPermissons) > 0 {
		_ = setRedisCacheRolesByUserId(id, userPermissons)
	}

	return userPermissons, nil
}
