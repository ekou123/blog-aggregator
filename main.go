package main

import (
	"example.com/sql/commands"
	"example.com/sql/internal/config"
	"fmt"
	"os"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	stateStruct := commands.State{
		&cfg,
	}

	cliCommands := commands.Commands{
		map[string]func(*commands.State, commands.Command) error{},
	}

	err = cliCommands.Register("login", commands.HandlerLogin)
	if err != nil {
		panic(err)
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
