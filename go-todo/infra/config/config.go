package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	HTTP_PORT  string
	JWT_SECRET string
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// http port
	HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		return envError("HTTP_PORT")
	}

	// jwt
	JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		return envError("JWT_SECRET")
	}

	return nil
}

func envError(env string) error {
	errStr := fmt.Sprintf("Error loading environment %v", env)
	return errors.New(errStr)
}
