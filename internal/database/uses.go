package database

import (
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/database/models"
)

func InsertUseCase(o *models.UseCase) bool {

	if len(o.UUID) == 0 {

		log.Error("Tried to insert Order with no uuid. This should not be possible.")

		return false
	}

	err := GetDB().Create(&o).Error

	if err != nil {

		log.Error(err)

		return false
	}

	return true
}
