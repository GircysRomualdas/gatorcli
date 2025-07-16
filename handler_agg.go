package main

import (
	"context"
	"fmt"
	"time"

	"github.com/GircysRomualdas/gatorcli/internal/database"
	"github.com/google/uuid"
)

const feedURL = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUserByName(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error: user %q not found\n", s.config.CurrentUserName)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error: failed to create feed: %w", err)
	}
	fmt.Printf("%+v\n", feed)

	return nil
}
