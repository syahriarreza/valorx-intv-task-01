package user

import "github.com/syahriarreza/valorx-intv-task-01/pkg/models"

type Usecase interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	Login(email, password string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error) // New method for OAuth
}
