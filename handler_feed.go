package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seanmoakes/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feedName> <feedURL>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]
	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("couldn't get user %s from db: %w", name, err)
	}
	user_id := usr.ID

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user_id,
	})
	if err != nil {
		return fmt.Errorf("unable to create feed %w", err)
	}

	fmt.Printf("Feed %s created successfully!\n", feed.Name)
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * URL:     %v\n", feed.Url)
	fmt.Printf(" * User ID: %v\n", feed.UserID)
}
