package main

import (
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read cfg file: %v\n", err)
	}
	programState := &state{cfg: &cfg}

	cmds := commands{cmds: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmd_name := os.Args[1]
	cmd_args := os.Args[2:]
	cmd := command{Name: cmd_name, Args: cmd_args}

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
