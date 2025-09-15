package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/har-sat/rssagg/internal/database"
	"github.com/har-sat/rssagg/internal/utils"
)

// threads -> no of goroutines that run simultaneously(i.e. the no of feeds being fetched simultaneously)
func startScrapping(db *database.Queries, threads int, timeBetweenRequests time.Duration) {
	log.Printf("Scraping on %v goroutines every %v duration", threads, timeBetweenRequests)

	ticker:= time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(threads))
		if err != nil {
			log.Printf("error fetching feeds from db: %v\n", err)
			continue
		}
		
		wg := sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Go(func() {scrape_feed(db, feed)})
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
		log.Printf("item found: %v", item.Title) 
	}
	
}