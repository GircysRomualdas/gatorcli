package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return write(*c)
}

func Read() (Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	cfg := Config{}
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homePath, configFileName), nil
}

func write(cfg Config) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(cfg); err != nil {
		return err
	}
	return nil
}
