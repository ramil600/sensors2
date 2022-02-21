package handlers

import (
	"github.com/ramil600/sensors2/foundation/web"
	"github.com/ramil600/sensors2/business/mid"


func API() *web.App {
	app := web.NewApp()
	app.Handle("/", mid.Logger(handlers.Myhandler))
	return app
}
