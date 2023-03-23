package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Sentence struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	Index     uint      `gorm:"not null"`
	Text      string    `gorm:"not null; size:1023"`
	Emotion   uint      `gorm:"not null"`
	NoteId    uuid.UUID `gorm:"not null"`
}

func (sentence *Sentence) BeforeCreate(*gorm.DB) (_ error) {
	sentence.ID = uuid.New()
	return
}
