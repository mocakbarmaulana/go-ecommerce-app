package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ServerPort       string
	Dsn              string
	AppSecret        string
	TwilioAccountSid string
	TwilioAuthToken  string
	TwilioFromNumber string
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

	dsn := os.Getenv("DSN")

	if len(dsn) == 0 {
		return AppConfig{}, errors.New("dsn is not set")
	}

	appSecret := os.Getenv("APP_SECRET")

	if len(appSecret) == 0 {
		return AppConfig{}, errors.New("app secret is not set")
	}

	return AppConfig{
		ServerPort:       httpPort,
		Dsn:              dsn,
		AppSecret:        appSecret,
		TwilioAccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioFromNumber: os.Getenv("TWILIO_FROM_NUMBER"),
	}, nil
}
