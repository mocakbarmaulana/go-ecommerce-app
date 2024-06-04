package helper

import (
	"errors"
	"github.com/go-ecommerce-app/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

type Auth struct {
	Secret string
}

func (a *Auth) CreateHashPassword(password string) (string, error) {

	if len(password) < 6 {
		log.Printf("Password must be at least 6 characters")
		return "", errors.New("password must be at least 6 characters")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error while hashing password: %v", err)
		return "", errors.New("error while hashing password")
	}

	return string(hashPassword), nil
}

func (a *Auth) GenerateToken(id uint, email string, role string) (string, error) {

	if id == 0 || email == "" || role == "" {
		log.Printf("Invalid user data")
		return "", errors.New("Required user data is missing")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.Secret))

	if err != nil {
		log.Printf("Error while generating token: %v", err)
		return "", errors.New("error while generating token")
	}

	return tokenString, nil
}

func (a *Auth) VerifyPassword(plainPassword string, hashPassword string) error {

	if len(plainPassword) < 6 {
		log.Printf("Password must be at least 6 characters")
		return errors.New("password must be at least 6 characters")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))

	if err != nil {
		log.Printf("Invalid password %v", err)
		return errors.New("invalid password")
	}

	return nil
}

func (a *Auth) VerifyToken(token string) (domain.User, error) {
	tokenArray := strings.Split(token, " ")

	if len(tokenArray) != 2 {
		return domain.User{}, errors.New("invalid format of token")
	}

	if tokenArray[0] != "Bearer" {
		return domain.User{}, errors.New("invalid format of token")
	}

	tokenString := tokenArray[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
		log.Printf("Invalid signing token: %v", err)
		return domain.User{}, errors.New("invalid signing token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token expired")
		}

		user := domain.User{
			ID:       uint(claims["user_id"].(float64)),
			Email:    claims["email"].(string),
			UserType: claims["role"].(string),
		}

		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a *Auth) Authorize(ctx *fiber.Ctx) error {

	return nil
}

func (a *Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {

	return domain.User{}
}
