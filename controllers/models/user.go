package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	Email     string
	Password  string
	Notes     []Note `gorm:"foreignKey:UserId"`
}
