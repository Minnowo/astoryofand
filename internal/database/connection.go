package database

import (
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type DBConfig struct {
	DatabasePath string
}

func DBInit(conf *DBConfig) {
	ldb, err := gorm.Open(sqlite.Open(conf.DatabasePath), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Database connection established")

	db = ldb

	log.Info("Running automigrate")
	db.AutoMigrate(&models.KeyValueStore{})
	db.AutoMigrate(&models.Settings{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.UseCase{})

}

func GetDB() *gorm.DB {
	return db
}
