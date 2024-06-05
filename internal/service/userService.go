package service

import (
	"errors"
	"fmt"
	"github.com/go-ecommerce-app/internal/domain"
	"github.com/go-ecommerce-app/internal/dto"
	"github.com/go-ecommerce-app/internal/helper"
	"github.com/go-ecommerce-app/internal/repository"
	"time"
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

func (us *UserService) isVerifiedUser(id uint) bool {
	user, err := us.Repo.FindUserById(id)

	return err == nil && user.Verified
}

func (us *UserService) GetVerificationCode(u domain.User) (int, error) {
	// if user already has a verification code, return it
	if us.isVerifiedUser(u.ID) {
		return 0, fmt.Errorf("user already verified")
	}

	// generate a new verification code
	code, err := us.Auth.GenerateCode()

	if err != nil {
		return 0, err
	}

	user := domain.User{
		Expire: time.Now().Add(time.Minute * 30),
		Code:   code,
	}

	_, err = us.Repo.UpdateUser(u.ID, user)

	if err != nil {
		return 0, errors.New("unable to update verification code to user")
	}

	// send sms

	return code, nil
}

func (us *UserService) VerifyCode(id uint, code int) error {
	if us.isVerifiedUser(id) {
		return errors.New("user already verified")
	}

	user, err := us.Repo.FindUserById(id)

	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("invalid verification code")
	}

	if user.Expire.Before(time.Now()) {
		return errors.New("verification code expired")
	}

	user.Verified = true

	_, err = us.Repo.UpdateUser(user.ID, user)

	return err
}
