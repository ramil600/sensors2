package user

import (
	"context"
	"fmt"
	"net/mail"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
