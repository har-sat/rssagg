package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/har-sat/rssagg/internal/database"
	"github.com/har-sat/rssagg/internal/utils"
)

func (apiCfg *apiConfig) handlerCreateUserFeed(w http.ResponseWriter, r *http.Request, u database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	var params parameters
	err := utils.DecodeJson(r, &params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Decoding JSON: %v", err))
		return
	}

	feed_follow, err := apiCfg.DB.CreateUserFeed(r.Context(), database.CreateUserFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    u.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating follow: %v", err))
		return
	}
	
	respondWithJson(w, 201, feed_follow)
}

// func (apiCfg *apiConfig) handlerGetUserFeeds(w http.ResponseWriter, r *http.Request, u database.User) {
// 	feeds, err := apiCfg.DB.GetUserFeeds(u)
// 	if err != nil {

// 	}
// }
