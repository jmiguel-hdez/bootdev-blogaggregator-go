package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.Name)
	}
	username := cmd.Args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("unable to set username")
	}
	fmt.Printf("User has been set to %s\n", username)
	return nil
}
