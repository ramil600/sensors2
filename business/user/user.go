package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/ramil600/sensors2/business/user/db"
	"golang.org/x/crypto/bcrypt"
)

type Core struct {
	store db.Store
}

func NewCore(dbconn *sqlx.DB) Core {
	return Core{
		store: db.NewStore(dbconn),
	}
}

// Create core user given update time and NewUser struct with fields to update, returns updated db.User
func (c Core) Create(ctx context.Context, nu NewUser, now time.Time) (db.User, error) {

	if nu.Password != nu.PasswordConfirm {
		return db.User{}, fmt.Errorf("passwords do not match")
	}

	if _, err := mail.ParseAddress(nu.Email); err != nil {
		return db.User{}, fmt.Errorf("email validation failed: %s", err)
	}

	pwdhash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.User{}, err
	}

	dbUsr := db.User{
		ID:           uuid.NewString(),
		Name:         nu.Name,
		Email:        nu.Email,
		Roles:        nu.Roles,
		PasswordHash: pwdhash,
		DateCreated:  now,
		DateUpdated:  now,
	}

	err = c.store.Create(ctx, dbUsr)
	return dbUsr, err

}

// Update core user given user id, update time and updUsr struct with fields to update, returns updated db.User
func (c Core) Update(ctx context.Context, updUsr UserUpdate, user_id string, now time.Time) (db.User, error) {

	dbUsr, err := c.store.QueryById(ctx, user_id)
	log.Println(dbUsr)
	if err != nil {
		return db.User{}, err
	}
	if updUsr.Name != nil {
		dbUsr.Name = *updUsr.Name
	}
	if updUsr.Email != nil {
		if _, err := mail.ParseAddress(*updUsr.Email); err != nil {
			return db.User{}, fmt.Errorf("email validation failed: %s", err)
		}
		dbUsr.Email = *updUsr.Email
	}
	if updUsr.Roles != nil {
		dbUsr.Roles = make(pq.StringArray, len(updUsr.Roles))
		copy(dbUsr.Roles, updUsr.Roles)
	}
	if updUsr.Password != nil {
		if *updUsr.Password != *updUsr.PasswordConfirm {
			return db.User{}, errors.New("user passwords are not equal")
		}
		pwdhash, err := bcrypt.GenerateFromPassword([]byte(*updUsr.Password), bcrypt.DefaultCost)
		if err != nil {
			return db.User{}, err
		}
		dbUsr.PasswordHash = pwdhash

	}
	dbUsr.DateUpdated = now
	err = c.store.Update(ctx, dbUsr)
	if err != nil {
		return db.User{}, err
	}
	return dbUsr, nil

}

func (c Core) Delete(ctx context.Context, id string) error {

	if _, err := uuid.Parse(id); err != nil {
		return err
	}

	return c.store.Delete(ctx, id)

}

func (c Core) Return(ctx context.Context, id string) (User, error) {

	_, err := uuid.Parse(id)
	if err != nil {
		return User{}, err

	}

	dbusr, err := c.store.QueryById(ctx, id)
	if err != nil {
		return User{}, err
	}
	coreUser := User{
		Name:  dbusr.Name,
		Email: dbusr.Email,
		Roles: dbusr.Roles,
	}

	return coreUser, nil

}
