package api

import (
	"errors"
	"github.com/go-ecommerce-app/config"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func StartServer(config config.AppConfig) error {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	err := app.Listen(":" + config.ServerPort)

	if err != nil {
		return errors.New("error starting server")
	}

	return nil
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "I'm alive!"})
}
