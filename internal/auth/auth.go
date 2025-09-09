package auth

import (
	"errors"
	"net/http"
	"strings"
)


// Header example
// Authorzation: ApiKey {you_api_key}
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("auth key malformed")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of apiKey")
	}
	return vals[1], nil
}