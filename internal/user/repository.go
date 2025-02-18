package user

import (
	"github.com/google/uuid"
	"github.com/syahriarreza/valorx-intv-task-01/pkg/models"
)

type Repository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uuid.UUID) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
	GetUserByEmail(email string) (*models.User, error)
}
