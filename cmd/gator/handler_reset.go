package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to reset database: %w", err)

	}
	fmt.Println("Database was reset successfully")
	return nil
}
