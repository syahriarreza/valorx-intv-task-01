package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
