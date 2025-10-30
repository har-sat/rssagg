package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/har-sat/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	FeedID      uuid.UUID `json:"feed_id"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
}

func DatabaseUserToUser(usr database.User) User {
	return User{
		ID:        usr.ID,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		Name:      usr.Name,
		ApiKey:    usr.ApiKey,
	}
}

func DatabaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, 0, len(dbFeeds))
	for _, feed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToFeed(feed))
	}
	return feeds
}

func DatabasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		Title:       post.Name,
		Description: post.Description.String,
		PublishedAt: post.PublishedAt,
		FeedID:      post.FeedID,
		Url:         post.Url,
		CreatedAt:   post.CreatedAt,
	}
}

func DatabasePostsToPosts(posts []database.Post) []Post {
	res := make([]Post, 0, len(posts))
	for _, post := range posts {
		res = append(res, DatabasePostToPost(post))
	}
	return res
}
