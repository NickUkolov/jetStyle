package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Note struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	AuthorID  uint
	Content   string
	Public    bool
}

type NoteInput struct {
	Content string `json:"content"`
	Public  bool   `json:"public"`
}
