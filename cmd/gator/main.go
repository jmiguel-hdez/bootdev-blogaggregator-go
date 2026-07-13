package main

import _ "github.com/lib/pq"

import (
	"database/sql"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/config"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/database"
	"log"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read cfg file: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("unable to open database")
	}
	defer db.Close()

	dbQueries := database.New(db)
	programState := &state{cfg: &cfg, db: dbQueries}

	cmds := commands{cmds: make(map[string]func(*state, command) error)}
	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

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
