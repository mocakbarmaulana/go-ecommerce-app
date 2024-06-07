package rest

import (
	"github.com/go-ecommerce-app/config"
	"github.com/go-ecommerce-app/internal/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App       *fiber.App
	Db        *gorm.DB
	Auth      helper.Auth
	AppConfig config.AppConfig
}
