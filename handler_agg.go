package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
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

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func scrapeFeeds(s *state) error {
	feedRow, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to fetch next feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feedRow.ID,
		UpdatedAt: time.Now().UTC(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	})
	if err != nil {
		return fmt.Errorf("error marking as fetched: %w", err)
	}

	feed, err := fetchFeed(context.Background(), feedRow.Url)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %w", err)
	}
	printRSSFeed(feed)
	return nil

}
