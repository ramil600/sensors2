package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/ramil600/sensors2/app/user/handlers"
)

const (
	// Server Shutdown Deadline in ms
	SHUTDOWN_DEADLINE = 20
	SERVER_ADDR       = "0.0.0.0:8081"
)

//App will have logger and middleware in it
type App struct {
	mux *mux.Router
}

// Handle created for better readability and is a facade for app.mux.Handle
func (a App) Handle(path string, handler http.Handler) {
	a.mux.Handle(path, handler)
}

// ServeHTTP is implementation of a http.Handler interface
func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	a.mux.ServeHTTP(w, req)
}

// NewApp creates mux router wrapped in App struct, creates some routing
func NewApp() *App {
	app := App{
		mux: mux.NewRouter(),
	}

	app.Handle("/", http.HandlerFunc(handlers.Myhandler))
	app.Handle("/ramil", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		w.Write([]byte("Hello" + r.URL.Path))
	}))
	return &app
}

func main() {

	app := NewApp()

	//Construct your server here
	s := &http.Server{
		Addr:    SERVER_ADDR,
		Handler: app,
	}

	// Put server in the go routine so that we can catch error from it or signal
	// to terminate lower in select structure
	// Create buffered channel so goroutine exits without blocking, better for memory alloc
	chErr := make(chan error, 1)
	go func() {
		chErr <- s.ListenAndServe()
	}()

	// Listen to interrupt signal and exit if you encounter error
	// or the shutdown signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	//catch server exception or SIGINT/SIGTERM signals
	select {
	case err := <-chErr:
		log.Fatal(err)
	case <-sig:
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), SHUTDOWN_DEADLINE*time.Millisecond)
		defer cancel()

		// Shutdown server gracefully if deadline expires, do the hard close
		log.Println("Shutting down server gracefully..")
		if err := s.Shutdown(ctx); err != nil {
			log.Println("couldn't shutdown server gracefully, will do hard close:", err)
			s.Close()
		}

	}
	return
}
