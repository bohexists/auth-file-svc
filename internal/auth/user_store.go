package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bohexists/auth-file-svc/internal/models"
)

type UserStore struct {
	users []models.User
}

func NewUserStore(filePath string) (*UserStore, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open user file: %v", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read user file: %v", err)
	}

	var users []models.User
	if err := json.Unmarshal(byteValue, &users); err != nil {
		return nil, fmt.Errorf("failed to parse user file: %v", err)
	}

	return &UserStore{users: users}, nil
}

func (us *UserStore) GetUserByUsername(username string) (*models.User, error) {
	for _, user := range us.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
