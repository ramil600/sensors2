package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Decode(r *http.Request, val interface{}) error {
	dec := json.NewDecoder(r.Body)
	return dec.Decode(val)
}

func Params(r *http.Request, param string) string {
	params := mux.Vars(r)
	user_id, ok := params[param]
	if !ok {
		return ""
	}
	return user_id

}
