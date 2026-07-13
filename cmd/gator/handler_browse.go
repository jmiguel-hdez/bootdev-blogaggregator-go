package main

import (
	"context"
	"fmt"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
	"strconv"
)

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

	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %v ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %v\n", post.Url)
		fmt.Println("==============================================")
	}
	return nil
}
