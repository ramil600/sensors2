package web

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

//TraceID is custom type for adding the trace to context
type traceID int

// Handler is custom handler functions in our app that will handle http routes
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
	return &app
}

// Handle adds trace info and allows mux to route the traffic with modified context
func (a App) Handle(path string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request) {
		var mytrace traceID
		ctx := context.WithValue(r.Context(), mytrace, "we345-wder23-ewe32")
		handler(ctx, w, r)
	}
	a.mux.HandleFunc(path, h)
}
