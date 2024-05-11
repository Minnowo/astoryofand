package models

import (
	"time"

	"github.com/minnowo/astoryofand/internal/assets"
	"github.com/minnowo/astoryofand/internal/util"
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func ValidPassword(password string) bool {
	return !util.IsEmptyOrWhitespace(password) &&
		len(password) >= assets.PASSWORD_MIN_LEN &&
		len(password) <= assets.PASSWORD_MAX_LEN
}

func ValidPasswordbytes(password []byte) bool {
	return !util.BytesIsEmptyOrWhitespace(password) &&
		len(password) >= assets.PASSWORD_MIN_LEN &&
		len(password) <= assets.PASSWORD_MAX_LEN
}

func ValidUsername(username string) bool {
	return !util.IsEmptyOrWhitespace(username) &&
		len(username) >= assets.USERNAME_MIN_LEN &&
		len(username) <= assets.USERNAME_MAX_LEN
}
func (u *User) CheckValid() bool {
	return ValidUsername(u.Username) && ValidPassword(u.Password)
}

type TableUser struct {
	CreatedAt time.Time
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Username  string         `gorm:"primarykey"`
	Password  []byte
}
