package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var HTTP_PORT string

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// http port
	HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		return errors.New("Error loading environment HTTP_PORT")
	}

	return nil
}
