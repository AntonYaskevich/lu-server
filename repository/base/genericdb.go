package base

import (
	"github.com/jmcvetta/neoism"
)

func Transactional(db *neoism.Database, query *neoism.CypherQuery) error {
	tx, err := db.Begin([]*neoism.CypherQuery{query})

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil;
}
