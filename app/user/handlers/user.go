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
	Log  *log.Logger
}

func NewUser(db *sqlx.DB, log *log.Logger) User {
	usr := User{
		Core: user.NewCore(db),
		Log:  log,
	}
	return usr
}

func Myhandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (u User) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var nu user.NewUser

	if err := Decode(r, &nu); err != nil {
		u.Log.Fatal(err)
	}

	dbUser, err := u.Core.Create(ctx, nu, time.Now())
	if err != nil {
		u.Log.Println(err)
		Encode(w, dbUser, http.StatusInternalServerError)
		return
	}

	Encode(w, dbUser, http.StatusCreated)

}
