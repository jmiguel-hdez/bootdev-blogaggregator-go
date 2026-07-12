package main

import (
	"context"
	"fmt"
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
	fmt.Printf("Feed %s collected, %v posts found\n", feedRow.Name, len(feed.Channel.Item))
	printRSSFeed(feed)
	return nil

}
