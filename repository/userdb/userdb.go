package userdb

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/jmcvetta/neoism"
	"github.com/AntonYaskevich/lu-server/models"
	"github.com/AntonYaskevich/lu-server/repository/base"
	"log"
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
		Statement:`CREATE (node:User {username: {username}, password: {password}})
		 	   RETURN id(node) as id, node.username as username, node.password as password`,
		Parameters: neoism.Props{"username": user.Username, "password": pass},
		Result: &result,
	}
	dberr := base.Transactional(self.Database, &query)
	if dberr != nil {
		log.Println(dberr)
		return nil, dberr
	}
	return &result[0], nil;
}

func (self *UserDB) Read(id uint) (*models.User, error) {
	result := [] models.User{}
	query := neoism.CypherQuery{
		Statement:`MATCH (n:User)
		 	   WHERE id(n) = {id}
		 	   RETURN id(n) as id, n.username as username, n.password as password`,
		Parameters: neoism.Props{"id": id},
		Result: &result,
	}

	err := self.Database.Cypher(&query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result[0], nil;
}

func (self *UserDB) Update(user *models.User) (*models.User, error) {
	result := [] models.User{}
	query := neoism.CypherQuery{
		Statement:`MATCH (n:User)
		 	   WHERE id(n) = {id}
		 	   SET n.username = {username}
		 	   RETURN id(n) as id, n.username as username, n.password as password`,
		Parameters: neoism.Props{"id": user.Id, "username": user.Username},
		Result: &result,
	}

	dberr := base.Transactional(self.Database, &query)
	if dberr != nil {
		log.Println(dberr)
		return nil, dberr
	}
	return &result[0], nil;
}

func (self *UserDB) Delete(id uint) error {
	query := neoism.CypherQuery{
		Statement:`MATCH (n:User)
		 	   WHERE id(n) = {id}
		 	   DELETE n`,
		Parameters: neoism.Props{"id": id},
	}

	return base.Transactional(self.Database, &query);
}

