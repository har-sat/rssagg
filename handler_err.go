package main

import (
	"net/http"

	"github.com/har-sat/rssagg/internal/utils"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 500, "HardCoded Error")
}