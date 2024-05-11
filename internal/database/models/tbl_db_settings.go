package models

import (
	"time"

	"gorm.io/gorm"
)

type TableSettings struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Profile string `gorm:"primarykey"`

	BoxSetPrice float32
	StickerCost float32

	PublicKey string
}
