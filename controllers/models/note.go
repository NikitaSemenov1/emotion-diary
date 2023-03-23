package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Note struct {
	ID        uuid.UUID  `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	Emotion   uint       `gorm:"not null; default:0" json:"emotion"`
	UserId    uuid.UUID  `gorm:"not null" json:"userId"`
	Sentences []Sentence `gorm:"foreignKey:NoteId; constraint:OnDelete:CASCADE" json:"sentences"`
}

func (note *Note) BeforeCreate(*gorm.DB) (_ error) {
	note.ID = uuid.New()
	return
}
