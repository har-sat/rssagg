package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseHandler(w, 200, "OK")
}