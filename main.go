package main

import (
	"fmt"
	"github.com/jmiguel-hdez/bootdev-blogaggregator-go/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read cfg file: %v\n", err)
	}
	err = cfg.SetUser("orsted")
	if err != nil {
		log.Fatalf("unable to write cfg file: %v\n", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("unable to read cfg file: %v\n", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

}
