package api

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func RetrieveAuthToken() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		return "", fmt.Errorf("TOKEN not found in .env file")
	}

	return token, nil
}