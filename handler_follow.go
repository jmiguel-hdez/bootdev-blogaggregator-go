package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
	"strconv"
	"time"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

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

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("Usage: %s", cmd.Name)
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

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	var err error
	if len(cmd.Args) == 1 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("usage: %s <limit>", cmd.Name)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error getting posts %w", err)
	}

	fmt.Printf("Found %v Posts for User %v\n", len(posts), user.Name)
	fmt.Println("===================================================")

	for i, post := range posts {
		fmt.Printf("Post[%v] Title: %v\n", i, post.Title)
		fmt.Printf("Post[%v] Url: %v\n", i, post.Url)
		fmt.Println()
	}
	return nil
}
