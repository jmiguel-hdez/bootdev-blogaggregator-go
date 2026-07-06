package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <Name> <Url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	curUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Unable to get current user in db %w", err)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    curUser.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Unable to create feed: %w", err)
	}
	fmt.Println("Feed created succesfully")
	printFeed(feed)
	fmt.Println()
	fmt.Println("========================================")

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error when getting the feeds: %w", err)
	}
	for _, feed := range feeds {
		username, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error when getting user: %w", err)
		}
		printFeed(feed)
		fmt.Printf(" * UserName: %v\n", username)
	}
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID: %v\n", feed.ID)
	fmt.Printf(" * Created At: %v\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt: %v\n", feed.UpdatedAt)
	fmt.Printf(" * Name: %v\n", feed.Name)
	fmt.Printf(" * Url: %v\n", feed.Url)
	fmt.Printf(" * UserID: %v\n", feed.UserID)
}
