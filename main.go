package main

import (
	"fmt"
	"log"

	"github.com/orest-srbn/go-camp/internal/rss"
)

func main() {
	// RSS feed URL
	feedURL := "http://podcast.dou.ua/rss"

	// Download and parse RSS feed
	feed, err := rss.ParseFeed(feedURL)
	if err != nil {
		log.Fatalf("Error parsing RSS feed: %v", err)
	}

	// Print results
	fmt.Printf("Title: %s\n", feed.Title)
	fmt.Printf("Description: %s\n", feed.Description)
	fmt.Printf("Items: %d\n\n", len(feed.Items))

	for i, item := range feed.Items {
		fmt.Printf("%d. %s\n", i+1, item.Title)
		fmt.Printf("   Link: %s\n", item.Link)
		fmt.Printf("   Date: %s\n\n", item.PubDate)
	}
}
