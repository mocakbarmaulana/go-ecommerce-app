package api

import (
	"errors"
	"github.com/go-ecommerce-app/config"
	"github.com/go-ecommerce-app/internal/api/rest"
	"github.com/go-ecommerce-app/internal/api/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) error {
	app := fiber.New()

	restHandler := &rest.RestHandler{
		App: app,
	}

	setupRoutes(restHandler)

	err := app.Listen(":" + config.ServerPort)

	if err != nil {
		return errors.New("error starting server")
	}

	return nil
}

func setupRoutes(rh *rest.RestHandler) {
	// user handlers
	handlers.SetupUserRoutes(rh)
	// transaction handlers
	// product handlers
}
