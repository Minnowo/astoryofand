package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/minnowo/astoryofand/internal/util"
	"gorm.io/gorm"
)

type UseCase struct {
	CreatedAt time.Time
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UUID      string         `gorm:"primarykey" json:"uuid"`
	UserData
	Email    string `json:"email" form:"email"`
	FullName string `json:"fullname" form:"fullname"`
	Comment  string `json:"comment" form:"comment"`
}

func NewUseCase() *UseCase {
	return &UseCase{
		UserData: UserData{UsecaseType},
	}
}

func (o *UseCase) DelayedInit() *UseCase {
	if err := uuid.Validate(o.UUID); err != nil {

		o.UUID = util.GetOrderID()
	}

	o.CreatedAt = time.Now()
	o.UpdatedAt = o.CreatedAt

	return o
}

func (o *UseCase) EnsureType() *UseCase {
	o.Type = UsecaseType
	return o
}

func (o *UseCase) CheckValidDataFromUser() bool {

	if util.IsEmptyOrWhitespace(o.Email) ||
		util.IsEmptyOrWhitespace(o.FullName) ||
		util.IsEmptyOrWhitespace(o.Comment) {
		return false
	}

	return true
}
