package service

import (
	"errors"
	"fmt"
	"github.com/go-ecommerce-app/internal/domain"
	"github.com/go-ecommerce-app/internal/dto"
	"github.com/go-ecommerce-app/internal/repository"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func (us *UserService) Register(input dto.UserRegister) (string, error) {

	user, err := us.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
	})

	// generate Token
	log.Println(user)

	userInfo := fmt.Sprintf("%v, %v, %v", user.Email, user.Phone, user.UserType)

	return userInfo, err
}

func (us *UserService) LoginUser(input dto.UserLogin) (string, error) {
	user, err := us.findUserByEmail(input.Email)

	if err != nil {
		return "", errors.New("user does not exist with this email")
	}

	return user.Email, nil
}

func (us *UserService) findUserByEmail(email string) (*domain.User, error) {

	user, err := us.Repo.FindUserByEmail(email)

	return &user, err
}
