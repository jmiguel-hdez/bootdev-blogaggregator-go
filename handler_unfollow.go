package main

import (
	"context"
	"fmt"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <Url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	s.db.DeleteUserFeedFollow(context.Background(), database.DeleteUserFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("error during delete query %w", err)
	}

	return nil
}
