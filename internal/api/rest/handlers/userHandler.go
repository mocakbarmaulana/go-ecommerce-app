package handlers

import (
	"github.com/go-ecommerce-app/internal/api/rest"
	"github.com/go-ecommerce-app/internal/dto"
	"github.com/go-ecommerce-app/internal/repository"
	"github.com/go-ecommerce-app/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.UserService{
		Repo:      repository.NewUserRepository(rh.Db),
		Auth:      rh.Auth,
		AppConfig: rh.AppConfig,
	}
	handler := UserHandler{
		svc: svc,
	}

	apiRoute := app.Group("/api/v1")

	pubApiRoute := apiRoute.Group("/auth")
	pubApiRoute.Post("/register", handler.Register)
	pubApiRoute.Post("/login", handler.Login)

	// private routes
	priApiRoute := apiRoute.Group("/users", rh.Auth.Authorize)
	priApiRoute.Get("/profile", handler.GetUser)
	priApiRoute.Get("/verification-code", handler.GetVerificationCode)
	priApiRoute.Post("/verify", handler.VerifyUser)
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
		"message": "Register",
		"token":   token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide valid input",
		})
	}

	token, err := h.svc.LoginUser(loginInput)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Please provide valid credentials",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login",
		"token":   token,
	})
}

func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	usr := h.svc.Auth.GetCurrentUser(ctx)

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User retrieved",
		"data":    usr,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	// create verification code and update user
	code, err := h.svc.GetVerificationCode(user)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get verification code",
			"error":   err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success",
		"code":    code,
	})
}

func (h *UserHandler) VerifyUser(ctx *fiber.Ctx) error {
	usr := h.svc.Auth.GetCurrentUser(ctx)

	var request = dto.VerifyUser{}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Please provide valid input",
		})
	}

	err := h.svc.VerifyCode(usr.ID, request.Code)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to verify user",
			"error":   err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User successfully verified",
	})
}
