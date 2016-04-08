package models

type User struct {
	Id       uint `json:"id(node)"`
	Username string `json:"node.username"`
	Password string `json:"node.password"`
}
