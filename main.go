package main

import (
	"github.com/jmcvetta/neoism"

	"github.com/AntonYaskevich/lu-server/repository/userdb"
	"github.com/AntonYaskevich/lu-server/models"
	"fmt"
)

func main() {

	db, err := neoism.Connect("http://neo4j:admin@localhost:7474/db/data")
	userdao := userdb.UserDB{db}
	user, err := userdao.Create(&models.User{Username:"white", Password:"some"})
	fmt.Println(user)
	user, err = userdao.Read(user.Id)
	fmt.Println(user)
	user.Username = "nigga"
	user, err = userdao.Update(user)
	fmt.Println(user)
	if err != nil {
		fmt.Println(err)
	}
}