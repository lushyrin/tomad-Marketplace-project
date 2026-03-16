package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.NewString()
	return nil
}
