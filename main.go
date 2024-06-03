package main

import (
	"github.com/go-ecommerce-app/config"
	"github.com/go-ecommerce-app/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("error loading env: %v\n", err)
	}

	err = api.StartServer(cfg)

	if err != nil {
		log.Fatalf("error starting server: %v\n", err)
	}
}
