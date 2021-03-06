package handlers

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/ramil600/sensors2/foundation/web"
)

func API(db *sqlx.DB, log *log.Logger) *web.App {
	app := web.NewApp(web.WithLogger(log))

	user := NewUser(db, log)
	app.Handle("/user/create", user.Create)
	app.Handle("/user/update/{id}", user.Update)
	app.Handle("/user/{id}", user.UserReturn)
	return app
}
