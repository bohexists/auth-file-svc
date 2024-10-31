// cmd/main.go
package main

import (
	"fmt"
	"log"

	"github.com/bohexists/auth-file-svc/config"
	"github.com/bohexists/auth-file-svc/internal/auth"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	userStore, err := auth.NewUserStore(cfg.UserFilePath)
	if err != nil {
		log.Fatalf("failed to load user store: %v", err)
	}

	authService := auth.NewAuthService(userStore)

	username := "exampleUser"
	password := "examplePassword"

	user, err := authService.Authenticate(username, password)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
	} else {
		fmt.Printf("User %s authenticated successfully with role %s\n", user.Username, user.Role)
	}
}
