package handlers

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, val interface{}) error {
	dec := json.NewDecoder(r.Body)
	return dec.Decode(val)
}
