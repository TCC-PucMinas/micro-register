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
	keyTokenRedisUserByCodeForgot       = "key-user-by-forgot-code"
	keyTokenRedisUserByCodeActive       = "key-user-by-active-code"
	keyTokenRedisUserByEmailAndCpf_Cpnj = "key-user-by-email-and-cpf_cnpj"
)

type User struct {
	Id         int64     `json:"id"`
	Phone      string    `json:"phone"`
	Forgot     string    `json:"forgot"`
	Business   string    `json:"business"`
	CpfCnpj    string    `json:"cpf_cnpj"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	CodeActive string    `json:"code_active`
	Active     int       `json:"active"`
	Adress     []Address `json:"addresses"`
}

func setRedisCacheUserByForgotCode(user *User) error {
	db := db.ConnectDatabaseRedis()

	json, err := json.Marshal(user)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyTokenRedisUserByCodeForgot, user.Forgot)

	return db.Set(key, json, 1*time.Hour).Err()
}

func getRedisCacheUserByForgotCode(forgot string) (User, error) {
	user := User{}

	redis := db.ConnectDatabaseRedis()

	key := fmt.Sprintf("%v - %v", keyTokenRedisUserByCodeForgot, forgot)

	value, err := redis.Get(key).Result()

	if err != nil {
		return user, err
	}

	if err := json.Unmarshal([]byte(value), &user); err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) GetOneUserByForgotCode() error {

	if user, err := getRedisCacheUserByForgotCode(u.Forgot); err == nil {
		u = &user
		return nil
	}

	sql := db.ConnectDatabase()

	query := `select id  from users where forgot = ? limit 1;`

	requestConfig, err := sql.Query(query, u.Forgot)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var id string
		_ = requestConfig.Scan(&id)
		i64, _ := strconv.ParseInt(id, 10, 64)
		u.Id = i64
	}

	if u.Id == 0 {
		return errors.New("Not found key")
	}

	_ = setRedisCacheUserByForgotCode(u)

	return nil
}

func (u *User) UpdateUserForgotById() (bool, error) {
	sql := db.ConnectDatabase()

	query := "update users set forgot = ? where id = ?"

	updateHealthyCheck, err := sql.Prepare(query)

	if err != nil {
		return false, err
	}

	_, e := updateHealthyCheck.Exec(u.Forgot, u.Id)

	if e != nil {
		return false, e
	}

	_ = setRedisCacheUserByForgotCode(u)

	return true, nil
}

func setRedisCacheUserByCodeActive(user *User) error {
	db := db.ConnectDatabaseRedis()

	json, err := json.Marshal(user)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v - %v", keyTokenRedisUserByCodeActive, user.CodeActive, user.Active)

	return db.Set(key, json, 1*time.Hour).Err()
}

func getRedisCacheUserByCodeActive(u *User) (User, error) {
	user := User{}

	redis := db.ConnectDatabaseRedis()

	key := fmt.Sprintf("%v - %v -%v", keyTokenRedisUserByCodeActive, u.CodeActive, u.Active)

	value, err := redis.Get(key).Result()

	if err != nil {
		return user, err
	}

	if err := json.Unmarshal([]byte(value), &user); err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) GetOneUserByCodeActive() error {

	if user, err := getRedisCacheUserByCodeActive(u); err == nil {
		u = &user
		return nil
	}

	sql := db.ConnectDatabase()

	query := `select id from users where code_active = ? and active = ? limit 1;`

	user, err := sql.Query(query, u.CodeActive, u.Active)

	if err != nil {
		return err
	}

	for user.Next() {
		var id string
		_ = user.Scan(&id)
		i64, _ := strconv.ParseInt(id, 10, 64)
		u.Id = i64
	}

	if u.Id == 0 {
		return errors.New("Not found key")
	}

	return nil
}

func (u *User) UpdateUserSetActiveUser() error {
	sql := db.ConnectDatabase()

	query := "update users set code_active = ?, active = ? where code_active = ?"

	updateUser, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := updateUser.Exec(nil, u.Active, u.CodeActive)

	if e != nil {
		return e
	}

	key := fmt.Sprintf("%v - %v - %v", keyTokenRedisUserByCodeActive, u.CodeActive, 1)

	db.RemoveCacheRedisByKey(key)

	return nil
}

func (u *User) UpdateUserSetNewPasswordByForgot() error {
	sql := db.ConnectDatabase()

	query := "update users set password = ?, forgot = ? where forgot = ?"

	updateUser, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := updateUser.Exec(u.Password, nil, u.Forgot)

	if e != nil {
		return e
	}

	return nil
}

func (u *User) CreateUser() (int64, error) {
	sql := db.ConnectDatabase()

	query := `insert into users (phone, business, cpf_cnpj, email, password, first_name, last_name, code_active, active)  values (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	insertUser, err := sql.Prepare(query)

	if err != nil {
		return 0, err
	}

	d, e := insertUser.Exec(u.Phone, u.Business, u.CpfCnpj, u.Email, u.Password, u.FirstName, u.LastName, u.CodeActive, u.Active)

	if e != nil {
		return 0, e
	}

	return d.LastInsertId()
}

func setRedisCacheUserByEmailAndCpf_Cpnj(user *User) error {
	db := db.ConnectDatabaseRedis()

	json, err := json.Marshal(user)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v - %v", keyTokenRedisUserByEmailAndCpf_Cpnj, user.Email, user.CpfCnpj)

	return db.Set(key, json, 1*time.Hour).Err()
}

func getRedisCacheUserByEmailAndCpf_Cpnj(u *User) (User, error) {
	user := User{}

	redis := db.ConnectDatabaseRedis()

	key := fmt.Sprintf("%v - %v -%v", keyTokenRedisUserByEmailAndCpf_Cpnj, user.Email, user.CpfCnpj)

	value, err := redis.Get(key).Result()

	if err != nil {
		return user, err
	}

	if err := json.Unmarshal([]byte(value), &user); err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) GetOneUserByEmailAndCpf_Cpnj() error {

	if user, err := getRedisCacheUserByEmailAndCpf_Cpnj(u); err == nil {
		u = &user
		return nil
	}

	sql := db.ConnectDatabase()

	query := `select id from users where email = ? or cpf_cnpj = ? limit 1;`

	user, err := sql.Query(query, u.Email, u.CpfCnpj)

	if err != nil {
		return err
	}

	for user.Next() {
		var id string
		_ = user.Scan(&id)
		i64, _ := strconv.ParseInt(id, 10, 64)
		u.Id = i64
	}

	if u.Id == 0 {
		return errors.New("Not found key")
	}

	_ = setRedisCacheUserByEmailAndCpf_Cpnj(u)

	return nil
}
