package handlers

import (
	"github.com/ramil600/sensors2/foundation/web"
)

func API() *web.App {
	app := web.NewApp(web.Logger)
	app.Handle("/", Myhandler)
	return app
}
