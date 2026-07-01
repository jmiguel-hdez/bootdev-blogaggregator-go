package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

// This method runs a given command with the provided state if it exists
func (c *commands) run(s *state, cmd command) error {
	f, ok := c.cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("%s is not a registered command", cmd.Name)
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

// This method registers a new handler function for a command name
func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}
