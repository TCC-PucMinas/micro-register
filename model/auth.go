package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/TCC-PucMinas/micro-register/db"
)

var (
	keyTokenRedisAuthByEmail = "key-user-auth-email"
)

type Authenticate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func setRedisCacheAuthenticate(user User) error {
	db := db.ConnectDatabaseRedis()

	json, err := json.Marshal(user)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyTokenRedisAuthByEmail, json)

	return db.Set(key, json, 1*time.Hour).Err()
}

func getRedisCacheAuthenticate(email string) (User, error) {
	user := User{}

	redis := db.ConnectDatabaseRedis()

	key := fmt.Sprintf("%v - %v", keyTokenRedisAuthByEmail, email)

	value, err := redis.Get(key).Result()

	if err != nil {
		return user, err
	}

	if err := json.Unmarshal([]byte(value), &user); err != nil {
		return user, err
	}

	return user, nil
}

func (auth *Authenticate) GetOneUserByEmail() (User, error) {

	if user, err := getRedisCacheAuthenticate(auth.Email); err == nil {
		return user, err
	}

	sql := db.ConnectDatabase()

	user := User{}

	query := `select id, first_name, last_name, email, phone, business, password from users where email = ? limit 1;`

	requestConfig, err := sql.Query(query, auth.Email)

	if err != nil {
		return user, err
	}

	for requestConfig.Next() {
		var id, firstName, lastName, email, phone, bussines, password string
		_ = requestConfig.Scan(&id, &firstName, &lastName, &email, &phone, &bussines, &password)
		i64, _ := strconv.ParseInt(id, 10, 64)
		user.Id = i64
		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.Phone = phone
		user.Business = bussines
		user.Password = password
	}

	if user.Id == 0 {
		return user, errors.New("User not found!")
	}

	_ = setRedisCacheAuthenticate(user)

	return user, nil
}
