package handlers

import (
	"fmt"
	"net/http"
)

func Myhandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		for _, h := range v {
			fmt.Println(k, h)
		}

	}
}
