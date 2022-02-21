package web

import (
	"context"
	"fmt"
	"net/http"
)

// Logger will implement some logging functionality
func Logger(h Handler) Handler {

	next := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		var ctxid traceID
		trace, ok := ctx.Value(ctxid).(string)
		if !ok {
			fmt.Println("No trace id found")
		}
		fmt.Println("Trace detected: ", trace)

		h(ctx, w, r)

	}
	return next
}
