package database

import "github.com/jmoiron/sqlx"

func Transact(db *sqlx.DB, fn func(sqlx.ExtContext) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			// rethrow panic after rolling back
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()
	err = fn(tx)
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}
