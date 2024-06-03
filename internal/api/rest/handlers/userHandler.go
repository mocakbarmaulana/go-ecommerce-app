package handlers

import (
	"github.com/go-ecommerce-app/internal/api/rest"
	"github.com/go-ecommerce-app/internal/dto"
	"github.com/go-ecommerce-app/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.UserService{}
	handler := UserHandler{
		svc: svc,
	}

	// public routes
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// private routes
	app.Get("/profile", handler.GetUser)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserRegister{}
	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := h.svc.Register(user)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": token,
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
