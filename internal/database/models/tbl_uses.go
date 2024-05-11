package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/minnowo/astoryofand/internal/util"
	"gorm.io/gorm"
)

type TableUseCase struct {
	CreatedAt time.Time
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UUID      string         `gorm:"primarykey" json:"uuid"`
	UserData
	Email    string `json:"email" form:"email"`
	FullName string `json:"fullname" form:"fullname"`
	Comment  string `json:"comment" form:"comment"`
}

func NewUseCase() *TableUseCase {
	return &TableUseCase{
		UserData: UserData{UsecaseType},
	}
}

func (o *TableUseCase) DelayedInit() *TableUseCase {
	if err := uuid.Validate(o.UUID); err != nil {

		o.UUID = util.GetOrderID()
	}

	o.CreatedAt = time.Now()
	o.UpdatedAt = o.CreatedAt

	return o
}

func (o *TableUseCase) EnsureType() *TableUseCase {
	o.Type = UsecaseType
	return o
}

func (o *TableUseCase) CheckValidDataFromUser() error {

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
