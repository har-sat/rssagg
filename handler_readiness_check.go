package main

import (
	"net/http"

	"github.com/har-sat/rssagg/internal/utils"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, "OK")
}