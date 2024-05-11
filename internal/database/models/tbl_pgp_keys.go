package models

import (
	"time"

	"gorm.io/gorm"
)

type TablePGPKey struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	PublicKey string         `gorm:"uniqueIndex"`
}
