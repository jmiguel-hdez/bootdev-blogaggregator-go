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
	fmt.Println("Feed follow created succesfully")
	printFollow(follow)
	fmt.Println()
	fmt.Println("======================================================")
	return nil
}

func printFollow(followRow database.CreateFeedFollowRow) {
	fmt.Printf(" * ID: %v\n", followRow.ID)
	fmt.Printf(" * CreatedAt: %v\n", followRow.CreatedAt)
	fmt.Printf(" * UpdatedAt: %v\n", followRow.UpdatedAt)
	fmt.Printf(" * UserID: %v\n", followRow.UserID)
	fmt.Printf(" * FeedID: %v\n", followRow.FeedID)
	fmt.Printf(" * FeedName: %v\n", followRow.FeedName)
	fmt.Printf(" * Username: %v\n", followRow.UserName)
}
