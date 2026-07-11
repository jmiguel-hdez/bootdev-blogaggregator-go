package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
	"time"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	//Need to get data for current user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error when getting current user data: %w", err)
	}

	//Need to get feed data for URL passed.
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error when getting feed: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed follow created succesfully")
	printFeedFollow(follow.UserName, follow.FeedName)
	fmt.Println()
	fmt.Println("======================================================")
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("Usage: %s", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user data: %w", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %w", err)
	}
	if len(feeds) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)

	printFeeds(feeds)
	fmt.Println()
	fmt.Println("===========================================================")

	return nil
}

func printFeeds(follows []database.GetFeedFollowsForUserRow) {
	for _, feed := range follows {
		fmt.Printf(" * Name: %v\n", feed.FeedName)
	}
}

func printFeedFollow(username, feedname string) {
	fmt.Printf(" * FeedName: %v\n", feedname)
	fmt.Printf(" * Username: %v\n", username)
}
