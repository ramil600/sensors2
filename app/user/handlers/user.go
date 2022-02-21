// hadndles package
package handlers

import (
	"context"
	"net/http"
)

func Myhandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
