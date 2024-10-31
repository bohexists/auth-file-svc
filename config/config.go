package config

import (
	"errors"
	"os"
)

type Config struct {
	UserFilePath string
}

func LoadConfig() (*Config, error) {
	userFilePath := os.Getenv("USER_FILE_PATH")
	if userFilePath == "" {
		userFilePath = "users.json"
		if _, err := os.Stat(userFilePath); errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("USER_FILE_PATH environment variable not set and default file not found")
		}
	}

	return &Config{UserFilePath: userFilePath}, nil
}
