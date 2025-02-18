package usecase

import (
	"errors"

	"github.com/syahriarreza/valorx-intv-task-01/internal/user"
	"github.com/syahriarreza/valorx-intv-task-01/pkg/models"
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
	return u.UserRepo.Create(user)
}

func (u *UserUsecase) GetUserByID(id string) (*models.User, error) {
	return u.UserRepo.GetByID(id)
}

func (u *UserUsecase) UpdateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	return u.UserRepo.Update(user)
}

func (u *UserUsecase) DeleteUser(id string) error {
	return u.UserRepo.Delete(id)
}
