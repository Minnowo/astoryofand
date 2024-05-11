package database

import (
	"fmt"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/database/models"
	"gorm.io/gorm"
)

var db *gorm.DB

type DBConfig interface {
	GetDSN() string
}

type SqliteDBConf struct {
	DatabasePath string
}

type PostgresDBConf struct {
	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabasePort     uint
	Others           map[string]interface{}
}

func DBInit(dialector gorm.Dialector, gconf *gorm.Config) {

	ldb, err := gorm.Open(dialector, gconf)

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

func (conf *SqliteDBConf) GetDSN() string {
	return conf.DatabasePath
}

func (conf *PostgresDBConf) GetDSN() string {

	var keyValueString []string

	keyValueString = append(keyValueString, fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		conf.DatabaseHost,
		conf.DatabaseUser,
		conf.DatabasePassword,
		conf.DatabaseName,
		conf.DatabasePort,
	))

	for key, value := range conf.Others {
		keyValueString = append(keyValueString, fmt.Sprintf("%s=%v", key, value))
	}

	return strings.Join(keyValueString, " ")
}

func GetDB() *gorm.DB {
	return db
}
