package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/orest-srbn/go-camp/internal/db"
	"github.com/orest-srbn/go-camp/internal/rss"
	"github.com/pressly/goose/v3"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	if err := db.Init(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Error setting dialect: %v", err)
	}

	if err := goose.Up(db.GetDB(), "migrations"); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	// RSS feed URL
	feedURL := "http://podcast.dou.ua/rss"

	// Download and parse RSS feed
	feed, err := rss.ParseFeed(feedURL)
	if err != nil {
		log.Fatalf("Error parsing RSS feed: %v", err)
	}

	ctx := context.Background()

	// Process each item
	for _, item := range feed.Items {
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing date for item %s: %v", item.Title, err)
			continue
		}

		article := &db.Article{
			GUID:        item.GUID,
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			PubDate:     pubDate,
		}

		if err := article.Save(ctx); err != nil {
			if err == db.ErrArticleExists {
				log.Printf("Article already exists: %s", article.Title)
				continue
			}
			log.Printf("Error saving article %s: %v", item.Title, err)
			continue
		}

		fmt.Printf("Saved article: %s\n", article.Title)
	}
}
