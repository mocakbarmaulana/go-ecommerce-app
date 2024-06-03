package service

import (
	"github.com/go-ecommerce-app/internal/domain"
	"github.com/go-ecommerce-app/internal/dto"
	"log"
)

type UserService struct {
}

func (us *UserService) Register(register dto.UserRegister) (string, error) {
	log.Printf("Registering user: %v\n", register)

	return "token-auth", nil
}

func (us *UserService) LoginUser(input interface{}) (string, error) {
	return "User logged in", nil
}

func (us *UserService) findUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}
