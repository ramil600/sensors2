package web

import (
	"context"
	"log"
	"net/http"
)

//WithLogger inserts main logger into the tracing
func WithLogger(log *log.Logger) Middleware {
	m := func(h Handler) Handler {
		next := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

			trace, ok := ctx.Value(Mytrace).(string)
			if !ok {
				log.Println("No trace id found")
			}
			log.Println("Trace started: ", trace, r.URL.Path, r.RemoteAddr)
			h(ctx, w, r)
			log.Println("Traced ended:  ", trace, r.URL.Path, r.RemoteAddr)
		}
		return next
	}
	return m
}
