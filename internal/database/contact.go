package database

import (
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/database/models"
)

func InsertContact(o *models.TableContact) bool {

	if len(o.UUID) == 0 {

		log.Error("Tried to insert Contact with no uuid. This should not be possible.")

		return false
	}

	err := GetDB().Create(&o).Error

	if err != nil {

		log.Error(err)

		return false
	}

	return true
}
