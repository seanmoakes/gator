package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seanmoakes/gator/internal/database"
)

func handlerAddFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}

	url := cmd.Args[0]

	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("couldn't get user %s from db: %w", s.cfg.CurrentUser, err)
	}
	user_id := usr.ID

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed for %s from db: %w", url, err)
	}
	feed_id := feed.ID

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user_id,
		FeedID:    feed_id,
	})
	if err != nil {
		return fmt.Errorf("unable to create feed follow: %w", err)
	}

	fmt.Println("New follow created:")
	printFeedFollow(usr.Name, feed.Name)
	fmt.Printf("* Created At:    %v\n", feed_follow.CreatedAt)

	return nil
}

func handlerGetFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("couldn't get user %s from db: %w", s.cfg.CurrentUser, err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), (user.ID))
	if err != nil {
		return fmt.Errorf("couldn't get feed follows for %s: %w", s.cfg.CurrentUser, err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No follows found.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	fmt.Println("=====================================")
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* Name:          %s\n", username)
	fmt.Printf("* Feed Name:     %s\n", feedname)
}
