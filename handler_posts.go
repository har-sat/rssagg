package main

import (
	"fmt"
	"net/http"

	"github.com/har-sat/rssagg/internal/database"
	"github.com/har-sat/rssagg/internal/utils"
)

func (apiCfg *apiConfig) handlerGetUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Limit int `json:"limit"`
	}
	var params parameters
	err := utils.DecodeJson(r, &params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error decoding json: %v", err))
	}

	data, err := apiCfg.DB.GetUserPosts(r.Context(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  int32(params.Limit),
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error fetching posts: %v", err))
	}

	utils.RespondWithJson(w, 200, data)
}
