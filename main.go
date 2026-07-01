package main

import (
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/config"
	"log"
	"os"
)

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
