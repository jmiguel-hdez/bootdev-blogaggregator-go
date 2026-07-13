package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
	"github.com/lib/pq"
	"github.com/lib/pq/pqerror"
	"time"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("incorrect time between reqs paramater: %w", err)
	}

	fmt.Printf("Collecting feeds every %v...\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			ticker.Stop()
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	feedRow, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to fetch next feed: %w", err)
	}

	_, err = s.db.MarkFeedFetched(context.Background(), feedRow.ID)
	if err != nil {
		return fmt.Errorf("error marking as fetched: %w", err)
	}

	feed, err := fetchFeed(context.Background(), feedRow.Url)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %w", err)
	}

	for i, item := range feed.Channel.Item {
		var pubtime sql.NullTime
		pubtime.Time, err = time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			pubtime.Valid = false
		} else {
			pubtime.Valid = true
		}
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: pubtime,
			FeedID:      feedRow.ID,
		})
		if pqErr := pq.As(err, pqerror.UniqueViolation); pqErr == nil {
			if err != nil {
				return fmt.Errorf("unable to create post: %w", err)
			}
			fmt.Printf("Created Post %v: Title: %v\n", i, post.Title)
		}
	}
	fmt.Printf("Feed %s collected, %v posts found\n", feedRow.Name, len(feed.Channel.Item))

	return nil
}
