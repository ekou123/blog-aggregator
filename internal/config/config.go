package config

import (
	"encoding/json"
	"example.com/sql/commands"
	"fmt"
	"os"
	"path/filepath"
)

const ConfigFileName = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	User  string `json:"current_user_name"`
}

func Read() (Config, error) {

	configPath, err := GetConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Error getting config file path: %v", err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %w", err)
	}

	return cfg, nil
}

func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error reading user home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, ConfigFileName)

	return configPath, nil
}

func (c *Config) SetUser(username string) error {
	if username == "" {
		return fmt.Errorf("username cannot be blank")
	}

	c.User = username

	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	configPath, err := GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %w", err)
	}

	file, err := os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func (c *Config) Save(filename string) error {
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	return os.WriteFile(filename, data, 0644)
}

func handlerLogin(s *commands.State, cmd commands.Command) error {
	if s.Cfg == nil {
		return fmt.Errorf("No config found")
	}

	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("No arguments passed to login")
	}

	username := cmd.Arguments[0]

	s.Cfg.User = username

	err := s.Cfg.Save(ConfigFileName)
	if err != nil {
		return err
	}

	fmt.Printf("User set to: %s\n", username)

	return nil
}
