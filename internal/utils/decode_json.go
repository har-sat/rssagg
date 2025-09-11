package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeJson(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}
