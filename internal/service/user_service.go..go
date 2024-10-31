package service

import (
	"encoding/json"
	"os"

	"github.com/bohexists/auth-file-svc/internal/models"
)

func AddUser(user models.User) error {
	filePath := os.Getenv("USER_FILE_PATH")

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func LoadUsers() ([]models.User, error) {
	filePath := os.Getenv("USER_FILE_PATH")

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []models.User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
