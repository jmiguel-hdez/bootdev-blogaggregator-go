package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
	"time"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <Name> <Url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Unable to create feed: %w", err)
	}
	paramsFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), paramsFeedFollow)
	if err != nil {
		return fmt.Errorf("unable to create feed follow: %w", err)
	}

	fmt.Println("Feed created succesfully")
	printFeed(feed, user)
	fmt.Println()
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("========================================")

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error when getting the feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error when getting user: %w", err)
		}
		printFeed(feed, user)
		fmt.Println("=============================================")
	}
	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf(" * ID: %v\n", feed.ID)
	fmt.Printf(" * Created At: %v\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt: %v\n", feed.UpdatedAt)
	fmt.Printf(" * Name: %v\n", feed.Name)
	fmt.Printf(" * Url: %v\n", feed.Url)
	fmt.Printf(" * User: %v\n", user.Name)
}
