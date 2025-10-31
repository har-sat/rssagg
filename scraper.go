package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/har-sat/rssagg/internal/database"
	"github.com/har-sat/rssagg/internal/utils"
)

// threads -> no of goroutines that run simultaneously(i.e. the no of feeds being fetched simultaneously)
func startScrapping(db *database.Queries, threads int, timeBetweenRequests time.Duration) {
	log.Printf("Scraping on %v goroutines every %v duration", threads, timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(threads))
		if err != nil {
			log.Printf("error fetching feeds from db: %v\n", err)
			continue
		}

		wg := sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Go(func() { scrape_feed(db, feed) })
		}
		wg.Wait()
	}
}

func scrape_feed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking feed with id %v fetched: %v", feed.ID, err)
		return
	}

	rssFeed, err := utils.UrlToFeed(feed.Url)
	if err != nil {
		log.Printf("error converting url %v to feed: %v", feed.Url, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		layouts := []string{
			time.RFC1123Z, 
			time.RFC1123,  
			"Mon, 02 Jan 2006 15:04:05 MST",
			"Mon, 2 Jan 2006 15:04:05 MST",
			"02 Jan 2006 15:04:05 MST",
		}

		var pubDate time.Time
		var parseErr error
		for _, layout := range layouts {
			pubDate, parseErr = time.Parse(layout, item.PubDate)
			if parseErr == nil {
				break
			}
		}
		if parseErr != nil {
			log.Printf("couldn't parse date %v with err: %v", item.PubDate, parseErr)
			continue
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			FeedID:    feed.ID,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: pubDate,
			Url:         item.Link,
		})
		if err != nil && !strings.Contains(err.Error(), "duplicate key") {
			log.Printf("couldn't create new post: %v", err)
			continue
		}
	}
	log.Printf("Feeds of %v scraped\n", feed.Name)
}
