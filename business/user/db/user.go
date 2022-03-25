package db

import (
	"context"
	"errors"
	"fmt"
	"reflect"

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
	defer rows.Close()

	if !rows.Next() {
		return User{}, errDbNotFound
	}
	err = rows.StructScan(&usr1)
	if err != nil {
		return User{}, errDbNotFound
	}

	return usr1, nil
}

// QuerySlice uses reflection to populate the slice of objects dest using
// query and struct with db tags
func (s Store) QuerySlice(ctx context.Context, query string, data interface{}, dest interface{}) error {
	val := reflect.ValueOf(dest)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		return errors.New("wrong data type slice is expected")
	}

	rows, err := s.DB.NamedQueryContext(ctx, query, data)
	if err != nil {
		return errors.New(fmt.Sprint("couldn't parse query", err))
	}
	defer rows.Close()
	// Get the slice from reflect pointer
	slice := val.Elem()

	for rows.Next() {
		// create new object of the type of the slice element
		elm := reflect.New(slice.Type().Elem())

		// reflect.Value extract interface
		if err = rows.StructScan(elm.Interface()); err != nil {
			return err
		}
		//use Set otherwise the interface object behind Value will not be updated
		slice.Set(reflect.Append(slice, elm.Elem()))

	}
	return nil

}

func (s Store) Delete(ctx context.Context, id string) error {
	const q = `DELETE FROM users WHERE "user_id"=:user_id`

	usrDel := struct {
		ID string `db:"user_id"`
	}{
		ID: id,
	}

	_, err := s.DB.NamedExecContext(ctx, q, usrDel)
	if err != nil {
		return err
	}
	return nil

}
