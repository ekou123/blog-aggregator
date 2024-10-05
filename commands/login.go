package commands

import (
	"example.com/sql/internal/config"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if s.Cfg == nil {
		return fmt.Errorf("No config found")
	}

	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("No arguments passed to login")
	}

	username := cmd.Arguments[0]

	s.Cfg.User = username

	err := s.Cfg.Save(config.ConfigFileName)
	if err != nil {
		return err
	}

	fmt.Printf("User set to: %s\n", username)

	return nil
}
