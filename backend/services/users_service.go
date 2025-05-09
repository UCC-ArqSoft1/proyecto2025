package services

import (
	"backend/clients"
	"fmt"
)

func Login(username string, password string) {
	// Get user by username from DB
	user := clients.GetUserByUsername(username)
	fmt.Println("Usuario obtenido", user)

	// Hash password

	// Compare hashes

	// Generate token
}
