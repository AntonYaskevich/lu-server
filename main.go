package main

import (
	"github.com/AntonYaskevich/lu-server/repository/userdb"
	"github.com/AntonYaskevich/lu-server/models"
	"fmt"
)

func main() {

	userdao := userdb.UserDB{}
	user, err1 := userdao.Create(&models.User{Username:"sample12", Password:"sample12"})
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(user)
}