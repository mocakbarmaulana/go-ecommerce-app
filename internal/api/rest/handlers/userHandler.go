package handlers

import (
	"github.com/go-ecommerce-app/internal/api/rest"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	handler := UserHandler{}

	// public routes
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// private routes
	app.Get("/profile", handler.GetUser)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User registered",
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User logged in",
	})
}

func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User retrieved",
	})
}
