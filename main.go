package main

import (
	"fmt"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

// This method runs a given command with the provided state if it exists
func (c *commands) run(s *state, cmd command) error {
	f, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("%s is not a registered command", cmd.name)
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Login expect one argument <username>")
	}
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("unable to set username")
	}
	fmt.Printf("User has been set to %s\n", username)
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read cfg file: %v\n", err)
	}
	s := state{}
	s.cfg = &cfg
	commands := commands{}
	commands.cmds = make(map[string]func(*state, command) error)
	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("You need to pass a command name")
	}
	cmd_name := os.Args[1]
	args := os.Args[2:]
	cmd := command{name: cmd_name, args: args}
	cmd_func, ok := commands.cmds[cmd_name]
	if !ok {
		log.Fatalf("cmd: %s is not a valid cmd", cmd_name)
	}
	err = cmd_func(&s, cmd)
	if err != nil {
		log.Fatalf("error when executing command %s, %v", cmd_name, err)
	}
}
