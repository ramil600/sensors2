package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ramil600/sensors2/app/user/handlers"
	"github.com/ramil600/sensors2/business/user/db"
	_ "github.com/ramil600/sensors2/foundation/web"
)

const (
	// Server Shutdown Deadline in ms
	SHUTDOWN_DEADLINE = 20
	SERVER_ADDR       = "0.0.0.0:8081"
)

/*
// Handle created for better readability and is a facade for app.mux.Handle
func (a App) Handle(path string, handler web.Handler) {
	a.mux.Handle(path, handler)
}
*/

func main() {

	// Establish DB connection from the config values
	db, err := db.Open(db.DBcfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("could not ping db conn: ", err)
	}

	//New logger setup
	log := log.New(os.Stdout, "USER: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	//Construct your server here
	s := &http.Server{
		Addr:    SERVER_ADDR,
		Handler: handlers.API(db, log), //custom built struct with mux, logger and middleware
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

}
