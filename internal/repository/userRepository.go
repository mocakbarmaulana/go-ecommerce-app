package repository

import (
	"errors"
	"github.com/go-ecommerce-app/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {
	err := r.db.Create(&usr).Error

	if err != nil {
		log.Printf("error creating user: %v\n", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return usr, nil
}

func (r userRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email = ?", email).Error

	if err != nil {
		log.Printf("error finding user by email: %v\n", err)
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, id).Error

	if err != nil {
		log.Printf("error finding user by id: %v\n", err)
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(u).Error

	if err != nil {
		log.Printf("error updating user: %v\n", err)
		return domain.User{}, errors.New("failed to update user")
	}

	return user, nil
}
