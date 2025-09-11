package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/har-sat/rssagg/internal/database"
	"github.com/har-sat/rssagg/internal/utils"
)

func (apiCfg *apiConfig) hanlderCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	var params parameters

	err := utils.DecodeJson(r, &params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Decoding json: %v", err))
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create feed: %v", err))
		return
	}

	respondWithJson(w, 201, DatabaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) hanlderGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't fetch feeds: %v", err))
		return
	}

	respondWithJson(w, 200, DatabaseFeedsToFeeds(feeds))
}
