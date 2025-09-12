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
		utils.RespondWithError(w, 400, fmt.Sprintf("Error Decoding JSON: %v", err))
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
		utils.RespondWithError(w, 400, fmt.Sprintf("Error Creating follow: %v", err))
		return
	}
	
	utils.RespondWithJson(w, 201, feed_follow)
}

func (apiCfg *apiConfig) handlerGetUserFeeds(w http.ResponseWriter, r *http.Request, u database.User) {
	feeds, err := apiCfg.DB.GetUserFeedFollows(r.Context(), u.ID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error getting user feed follows: %v", err))	
	}
	utils.RespondWithJson(w, 200, feeds)
}

func (apiCfg *apiConfig) handlerDeleteUserFeed(w http.ResponseWriter, r *http.Request, u database.User) {
	type parameters struct {
		FeedId uuid.UUID
	}
	var params parameters
	err := utils.DecodeJson(r, &params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error decoding JSON: %v", err))
	}

	err = apiCfg.DB.DeleteUserFeedFollow(r.Context(), database.DeleteUserFeedFollowParams{
		UserID: u.ID,
		FeedID: params.FeedId,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error deleting user feed follow: %v", err))	
	}
	utils.RespondWithJson(w, 200, fmt.Sprintf("Deleted feed with feed id: %v", params.FeedId))
}
