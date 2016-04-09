package models

type User struct {
	Id       uint `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
