package model

type Address struct {
	Id         int64  `json:"id"`
	Street     string `json:"street"`
	State      string `json:"state"`
	Number     string `json:"number"`
	Country    string `json:"country"`
	Complement string `json:"complement"`
}
