package main

import (
	"database/sql"
	"example.com/sql/commands"
	"example.com/sql/internal/config"
	"example.com/sql/internal/database"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

const dbURL = "postgres://postgres:postgres@localhost:5432/gator"

func main() {

	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	// Open database connection
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)

	stateStruct := config.State{
		Cfg: &cfg,
		Db:  dbQueries,
	}

	cliCommands := commands.Commands{
		map[string]func(*config.State, commands.Command) error{},
	}

	err = cliCommands.Register("login", commands.HandlerLogin)
	if err != nil {
		fmt.Println("Error registering command:", err)
		os.Exit(1)
	}

	args := os.Args

	commandName := args[1]
	commandArgs := args[2:]

	cmd := commands.Command{
		Name:      commandName,
		Arguments: commandArgs,
	}

	err = cliCommands.Run(&stateStruct, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

//type Commands struct {
//	CLICommands map[string]func(*State, Command) error
//}
