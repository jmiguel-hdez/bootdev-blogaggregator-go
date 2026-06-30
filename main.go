package main

import (
	"fmt"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("unable to read cfg file: %v\n", err)
	}
	err = config.SetUser(cfg, "orsted")
	if err != nil {
		fmt.Printf("unable to write cfg file: %v\n", err)
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("unable to read cfg file: %v\n", err)
	}
	fmt.Printf("username: %v\n", cfg.CurrentUserName)
	fmt.Printf("url: %v\n", cfg.DbUrl)

}
