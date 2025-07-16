package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/GircysRomualdas/gatorcli/internal/config"
	"github.com/GircysRomualdas/gatorcli/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func getState() (*state, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	}

	state := &state{
		config: &cfg,
	}

	return state, nil
}

func main() {
	programState, err := getState()
	if err != nil {
		log.Fatal(err)
	}

	cmds := getCommands()

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	db, err := sql.Open("postgres", programState.config.DbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	programState.db = dbQueries

	err = cmds.run(programState, command{
		Name: cmdName,
		Args: cmdArgs,
	})
	if err != nil {
		log.Fatal(err)
	}
}
