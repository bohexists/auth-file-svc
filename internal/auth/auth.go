package auth

import (
	"fmt"

	"github.com/bohexists/auth-file-svc/internal/models"
)

type AuthService struct {
	store *UserStore
}

func NewAuthService(store *UserStore) *AuthService {
	return &AuthService{store: store}
}

func (as *AuthService) Authenticate(username, password string) (*models.User, error) {
	user, err := as.store.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %v", err)
	}

	if user.Password != password {
		return nil, fmt.Errorf("authentication failed: incorrect password")
	}

	return user, nil
}
