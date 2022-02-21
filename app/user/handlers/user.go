// hadndles package
package handlers

import (
	"context"
	"fmt"
	"net/http"
)

func Myhandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		for _, h := range v {
			fmt.Println(k, h)
		}

	}
}
