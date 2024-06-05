package service

import (
	"errors"
	"github.com/go-ecommerce-app/internal/domain"
	"github.com/go-ecommerce-app/internal/dto"
	"github.com/go-ecommerce-app/internal/helper"
	"github.com/go-ecommerce-app/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (us *UserService) Register(input dto.UserRegister) (string, error) {

	hashPassword, err := us.Auth.CreateHashPassword(input.Password)

	if err != nil {
		return "", err
	}

	user, err := us.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Phone:    input.Phone,
		Password: hashPassword,
	})

	return us.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (us *UserService) LoginUser(input dto.UserLogin) (string, error) {
	user, err := us.findUserByEmail(input.Email)

	if err != nil {
		return "", errors.New("user does not exist with this email")
	}

	err = us.Auth.VerifyPassword(input.Password, user.Password)

	if err != nil {
		return "", err
	}

	token, err := us.Auth.GenerateToken(user.ID, user.Email, user.UserType)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) findUserByEmail(email string) (*domain.User, error) {

	user, err := us.Repo.FindUserByEmail(email)

	return &user, err
}
