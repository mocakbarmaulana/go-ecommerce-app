package api

import (
	"errors"
	"github.com/go-ecommerce-app/config"
	"github.com/go-ecommerce-app/internal/api/rest"
	"github.com/go-ecommerce-app/internal/api/rest/handlers"
	"github.com/go-ecommerce-app/internal/domain"
	"github.com/go-ecommerce-app/internal/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func StartServer(config config.AppConfig) error {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error connecting to db: %v\n", err)
	}

	log.Println("Database connected")

	// run migrations
	err = db.AutoMigrate(&domain.User{})

	if err != nil {
		log.Fatalf("error running migrations: %v\n", err)
	}

	log.Println("Migrations run successfully")

	auth := helper.SetupAuth(config.AppSecret)

	restHandler := &rest.RestHandler{
		App:       app,
		Db:        db,
		Auth:      auth,
		AppConfig: config,
	}

	setupRoutes(restHandler)

	err = app.Listen(":" + config.ServerPort)

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
