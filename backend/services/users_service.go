package services

import (
	"fmt"

	"backend/clients"
	"backend/utils"
)

func Login(username string, password string) (int, string, error) {
	userDAO, err := clients.GetUserByUsername(username)
	if err != nil {
		return 0, "", fmt.Errorf("error getting user: %w", err)
	}
	if utils.HashSHA256(password) != userDAO.PasswordHash {
		return 0, "", fmt.Errorf("invalid password")
	}
	token, err := utils.GenerateJWT(userDAO.ID)
	if err != nil {
		return 0, "", fmt.Errorf("error generating token: %w", err)
	}
	return userDAO.ID, token, nil
}
