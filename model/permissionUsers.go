package model

type UserPermission struct {
	Id         int64      `json:"id"`
	User       User       `json:"user"`
	Permission Permission `json:"permission"`
}
