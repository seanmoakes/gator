package main

import (
	"fmt"
	"log"

	"github.com/seanmoakes/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config: %v\n", cfg)

	err = cfg.SetUser("sean")
	if err != nil {
		log.Fatalf("error writing to config: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
