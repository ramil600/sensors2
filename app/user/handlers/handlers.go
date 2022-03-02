package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/ramil600/sensors2/foundation/web"
)

func API(db *sqlx.DB) *web.App {
	app := web.NewApp(web.Logger)

	user := NewUser(db)
	app.Handle("/user/create", user.Create)
	return app
}
