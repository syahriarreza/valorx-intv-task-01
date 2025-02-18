package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/syahriarreza/valorx-intv-task-01/internal/user"
	"github.com/syahriarreza/valorx-intv-task-01/pkg/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
