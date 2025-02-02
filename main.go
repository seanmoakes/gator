package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/seanmoakes/gator/internal/config"
	"github.com/seanmoakes/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	// Read existing config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Connect to database
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	// Set up program state
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	// register commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
