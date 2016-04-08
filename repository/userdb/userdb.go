package userdb

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/jmcvetta/neoism"
	"github.com/AntonYaskevich/lu-server/models"
)

type UserDB struct {
	Database *neoism.Database
}

func (self *UserDB) Create(user *models.User) (*models.User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	result := [] models.User{}
	query := neoism.CypherQuery{
		Statement:`CREATE (node:User {username: {username}, password: {password}}) RETURN id(node), node.username, node.password`,
		Parameters: neoism.Props{"username": user.Username, "password": pass},
		Result: &result,
	}
	self.Database.Cypher(&query)

	return &result[0], nil;
}