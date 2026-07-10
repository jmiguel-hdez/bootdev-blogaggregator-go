package main

import (
	"context"
	"fmt"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
)

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
	printFeeds(feeds)
	fmt.Println()
	fmt.Println("===========================================================")

	return nil
}

func printFeeds(follows []database.GetFeedFollowsForUserRow) {
	for _, feed := range follows {
		fmt.Printf(" * Name: %v\n", feed.FeedName)
		fmt.Printf(" * User %v\n", feed.UserName)
	}
}
