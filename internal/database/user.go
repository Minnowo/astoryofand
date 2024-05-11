package database

import (
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/database/models"
	"golang.org/x/crypto/bcrypt"
)

func InsertRawUser(u *models.User) bool {

	var usr models.TableUser

	usr.Username = u.Username
	usr.Password = []byte(u.Password)

	return InsertUser(&usr)
}

func InsertUser(o *models.TableUser) bool {

	if !models.ValidUsername(o.Username) {

		log.Debug("username is invalid")

		return false
	}

	if !models.ValidPasswordbytes(o.Password) {

		log.Debug("password is invalid")

		return false
	}

	password, err := bcrypt.GenerateFromPassword(o.Password, bcrypt.DefaultCost)

	if err != nil {

		log.Error(err)

		return false
	}

	o.Password = password

	err = GetDB().Create(&o).Error

	if err != nil {

		log.Error(err)

		return false
	}

	return true
}

func AuthUser(u *models.User) (bool, error) {

	if !u.CheckValid() {
		return false, nil
	}

	var user models.TableUser

	err := GetDB().First(&user, "username = ?", u.Username).Error

	if err != nil {

		log.Error(err)

		return false, nil
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(u.Password)); err != nil {

		return false, nil
	}

	return true, nil
}
