package web

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ramil600/sensors2/app/user/handlers"

	"github.com/ramil600/sensors2/business/mid"
)

//we will add trace for our application
type traceID int

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request)

//App will have logger and middleware in it
type App struct {
	mux *mux.Router
}

// ServeHTTP is implementation of a http.Handler interface for App
//It will allow it use http.Handler interface inside the server struct
func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	a.mux.ServeHTTP(w, req)
}

// NewApp creates mux router wrapped in App struct, creates some routing
func NewApp() *App {
	app := App{
		mux: mux.NewRouter(),
	}

	app.Handle("/", mid.Logger(handlers.Myhandler))
	/*
		app.Handle("/{name}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			w.Write([]byte("Hello, " + vars["name"]))
		}))

	*/
	return &app
}

// Handle created for better readability and is a facade for app.mux.Handle
func (a App) Handle(path string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request) {
		var mytrace traceID
		ctx := context.WithValue(r.Context(), mytrace, "we345-wder23-ewe32")

		handler(ctx, w, r)

	}

	a.mux.HandleFunc(path, h)
}
