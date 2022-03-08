package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var errDbNotFound = errors.New("not found")

// Store is encapsulation for db connection
type Store struct {
	DB *sqlx.DB
}

// NewStore creates new Store with db connection pool
func NewStore(db *sqlx.DB) Store {
	return Store{
		DB: db,
	}
}

// Create inserts a new user and logs the user_id created
func (s Store) Create(ctx context.Context, u User) error {
	u.ID = uuid.NewString()

	const query = `INSERT INTO users (user_id, name, email, roles, password_hash,
		date_created, date_updated) VALUES(:user_id,:name,:email,:roles,:password_hash,
			:date_created, :date_updated) RETURNING user_id`
	const queryid = `SELECT user_id FROM users WHERE email=$1`

	_, err := s.DB.NamedExecContext(ctx, query, u)
	if err != nil {
		return err
	}
	var userId string
	err = s.DB.QueryRowContext(ctx, queryid, u.Email).Scan(&userId)
	if err != nil {
		return err
	}
	return nil

}

func (s Store) Update(ctx context.Context, upd User) error {

	const q = `UPDATE users SET
	"name"=:name, 
	"email"=:email,
	"roles"=:roles,
	"password_hash"=:password_hash,
	"date_updated"=:date_updated
	WHERE "user_id"=:user_id`

	_, err := s.DB.NamedExecContext(ctx, q, upd)
	if err != nil {
		return err
	}
	return nil

}

func (s Store) QueryById(ctx context.Context, user_id string) (User, error) {
	const q = `SELECT * FROM users WHERE "user_id"=:user_id`

	var usr = struct {
		ID string `db:"user_id"`
	}{
		ID: user_id,
	}
	var usr1 User

	rows, err := s.DB.NamedQueryContext(ctx, q, usr)
	if err != nil {
		return User{}, errors.New(fmt.Sprint("couldn't parse query", err))
	}

	if !rows.Next() {
		return User{}, errDbNotFound
	}
	err = rows.StructScan(&usr1)
	if err != nil {
		return User{}, errDbNotFound
	}

	return usr1, nil
}
