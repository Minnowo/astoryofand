package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/minnowo/astoryofand/internal/util"
	"gorm.io/gorm"
)

type TableContact struct {
	CreatedAt time.Time
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UUID      string         `gorm:"primarykey" json:"uuid"`
	UserData
	Email    string `json:"email" form:"email"`
	FullName string `json:"fullname" form:"fullname"`
	Comment  string `json:"comment" form:"comment"`
}

func NewContact() *TableContact {
	return &TableContact{}
}

func (o *TableContact) EnsureType() *TableContact {
	o.Type = ContactUsType
	return o
}

func (o *TableContact) DelayedInit() *TableContact {

	if err := uuid.Validate(o.UUID); err != nil {

		o.UUID = uuid.NewString()
	}

	o.CreatedAt = time.Now()
	o.UpdatedAt = o.CreatedAt

	return o
}

func (o *TableContact) CheckValidDataFromUser() error {

	if util.IsEmptyOrWhitespace(o.Email) {
		return fmt.Errorf("Email is empty!")
	}

	if util.IsEmptyOrWhitespace(o.FullName) {
		return fmt.Errorf("FullName is empty!")
	}

	if util.IsEmptyOrWhitespace(o.Comment) {
		return fmt.Errorf("Comment is empty!")
	}

	return nil
}
