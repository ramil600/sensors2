package db

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	DB *sqlx.DB
}

// NewStore creates new Store with db connection pool
func NewStore(db *sqlx.DB) *Store {
	return &Store{
		DB: db,
	}
}

// Create inserts a new user and logs the user_id created
func (s Store) Create(ctx context.Context, nu NewUser) error {
	nu.ID = uuid.NewString()

	const query = `INSERT INTO users (user_id, name, email, roles, password_hash,
		date_created, date_updated) VALUES(:user_id,:name,:email,:roles,:password_hash,
			:date_created, :date_updated) RETURNING user_id`
	const queryid = `SELECT user_id FROM users WHERE email=$1`

	_, err := s.DB.NamedExecContext(ctx, query, nu)
	if err != nil {
		return err
	}
	var userId string
	err = s.DB.QueryRowContext(ctx, queryid, nu.Email).Scan(&userId)
	if err != nil {
		return err
	}
	log.Println("Id of created user: ", userId)

	return nil

}
