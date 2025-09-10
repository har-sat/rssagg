package main

import (
	"fmt"
	"net/http"

	"github.com/har-sat/rssagg/internal/auth"
	"github.com/har-sat/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("AuthError: %v", err))
			return
		}

		usr, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't Get User: %v", err))
			return
		}

		handler(w, r, usr)
	}
}
