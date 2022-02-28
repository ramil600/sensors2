// hadndles package
package handlers

import (
	"context"
	"net/http"

	"github.com/ramil600/sensors2/business/user/db"
)

type User struct {
	Store db.Store
}

func Myhandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (u User) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Insert Something"))
}
