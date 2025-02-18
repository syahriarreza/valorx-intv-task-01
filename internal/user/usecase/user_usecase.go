package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/syahriarreza/valorx-intv-task-01/internal/user"
	"github.com/syahriarreza/valorx-intv-task-01/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepo user.Repository
}

func NewUserUsecase(repo user.Repository) user.Usecase {
	return &UserUsecase{
		UserRepo: repo,
	}
}

func (u *UserUsecase) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	return u.UserRepo.CreateUser(user)
}

func (u *UserUsecase) GetUserByID(id string) (*models.User, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}
	return u.UserRepo.GetUserByID(parsedID)
}

func (u *UserUsecase) UpdateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	return u.UserRepo.UpdateUser(user)
}

func (u *UserUsecase) DeleteUser(id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid user ID format")
	}
	return u.UserRepo.DeleteUser(parsedID)
}

func (u *UserUsecase) Login(email, password string) (*models.User, error) {
	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

func (u *UserUsecase) GetUserByEmail(email string) (*models.User, error) {
	return u.UserRepo.GetUserByEmail(email)
}
