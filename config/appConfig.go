package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (cfg AppConfig, err error) {
	err = godotenv.Load()

	if err != nil {
		return AppConfig{}, err
	}

	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) == 0 {
		return AppConfig{}, errors.New("port is not set")
	}

	return AppConfig{ServerPort: httpPort}, nil
}
