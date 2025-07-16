package main

import (
	"context"
	"fmt"
	"time"

	"log"

	"github.com/GircysRomualdas/gatorcli/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.GetUserByName(context.Background(), name)
	if err != nil {
		log.Fatalf("error: user %q not found\n", name)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				log.Fatalf("error: user %q already exists\n", name)
			}
		}
		return fmt.Errorf("couldn't create user: %w", err)
	}
	fmt.Printf("User %s created successfully!\n", user.Name)
	s.config.SetUser(user.Name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error: couldn't reset users: %v", err)
	}
	fmt.Println("Users reset successfully!")
	return nil
}
