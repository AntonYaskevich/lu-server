package repository

import (
	"github.com/jmcvetta/neoism"
	"log"
)

const (
	Neo4jURL = "http://neo4j:admin@localhost:7474/db/data"
)

func New() *neoism.Database {
	db, err := neoism.Connect(Neo4jURL)
	if err != nil {
		log.Panic(err)
		return nil
	}
	return db
}
