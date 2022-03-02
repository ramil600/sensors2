// hadndles package
package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ramil600/sensors2/business/user"
)

type User struct {
	Core user.Core
}

func NewUser(db *sqlx.DB) User {
	usr := User{
		Core: user.NewCore(db),
	}
	return usr
}

func Myhandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (u User) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var nu user.NewUser

	if err := Decode(r, &nu); err != nil {
		log.Fatal(err)
	}

	dbUser, err := u.Core.Create(ctx, nu, time.Now())
	if err != nil {
		Encode(w, dbUser, http.StatusInternalServerError)
	}

	Encode(w, dbUser, http.StatusCreated)

}
