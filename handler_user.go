package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
	"time"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.Name)
	}
	username := cmd.Args[0]
	userid := uuid.New()
	ct := time.Now()
	params := database.CreateUserParams{ID: userid, CreatedAt: ct, UpdatedAt: ct, Name: username}

	_, err := s.db.GetUser(context.Background(), username)
	if err == nil {
		return fmt.Errorf("user:%s already exists", username)
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("User created: %+#v\n", user)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.Name)
	}
	username := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("can't login because user doesn't exist in db")
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("unable to set username")
	}
	fmt.Printf("User has been set to %s\n", username)
	return nil
}
