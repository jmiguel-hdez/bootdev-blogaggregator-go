package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(homedir, configFileName)

	return path, nil
}

func Read() (Config, error) {

	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func SetUser(cfg Config, current_user_name string) error {
	cfg.CurrentUserName = current_user_name

	return write(cfg)
}

func write(cfg Config) error {

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
