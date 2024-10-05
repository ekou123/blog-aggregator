package commands

import (
	"example.com/sql/internal/config"
	"fmt"
)

type State struct {
	Cfg *config.Config
}

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	CLICommands map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {

}

func (c *Commands) run(s *State, cmd Command) error {
	handler, ok := c.CLICommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command %s not found", cmd.Name)
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
