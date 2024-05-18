package usecase

import (
	"errors"

	"github.com/mokhlesurr031/goblog/internal/models"
	"github.com/mokhlesurr031/goblog/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	RegisterUser(user *models.User) error
	Login(email, password string) (*models.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) RegisterUser(user *models.User) error {
	// Check if the user already exists
	existingUser, _ := u.userRepo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already in use")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create the user
	return u.userRepo.CreateUser(user)
}

func (u *userUsecase) Login(email, password string) (*models.User, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
