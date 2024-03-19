package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/assets"
	"github.com/minnowo/astoryofand/internal/database"
	"github.com/minnowo/astoryofand/internal/database/models"
	"github.com/minnowo/astoryofand/internal/util"
)

type PGPDecryptor struct {
	PrivateKey string
	Password   []byte
}

var (
	Decryptor PGPDecryptor
)

func (pgp PGPDecryptor) processFile(path string) error {

	if filepath.Ext(path) != ".asc" {

		log.Warnf("Will not process %s because it does not end with .asc", path)

		return errors.New("Will not process")
	}

	log.Debug("Opening file...")

	fileContent, err := os.ReadFile(path)

	if err != nil {

		log.Error(err)

		return err
	}

	log.Debug("Decrypting file...")

	data, err := helper.DecryptBinaryMessageArmored(pgp.PrivateKey, pgp.Password, string(fileContent))

	if err != nil {

		log.Errorf("Could not decrpt file: %v", err)

		return err
	}

	log.Debug(data)

	log.Debug("Unmarshal file...")

	var e models.UserData

	if err := json.Unmarshal(data, &e); err != nil {

		log.Error(err)

		return err
	}

	log.Debug(e)

	switch e.Type {

	case models.OrderType:

		log.Info("Got Order from file")

		var o models.Order

		if err := json.Unmarshal(data, &o); err != nil {

			log.Error(err)

			return err
		}

		log.Debug(o)

        database.InsertOrder(&o)

		break

	case models.UsecaseType:

		log.Info("Got Usecase from file")

		var o models.UseCase

		if err := json.Unmarshal(data, &o); err != nil {

			log.Error(err)

			return err
		}

		log.Debug(o)
        database.InsertUseCase(&o)

		break
	}

	return nil
}

func watch(watcher *fsnotify.Watcher) {

	for {

		select {

		case event, ok := <-watcher.Events:

			if !ok {
				return
			}

			switch event.Op {

			case fsnotify.Create:
				log.Infof("File created: %s", event.Name)

				if err := Decryptor.processFile(event.Name); err == nil {
					log.Info("Success")
				}

				break

			case fsnotify.Write:
			case fsnotify.Remove:
			case fsnotify.Rename:
			case fsnotify.Chmod:
				break

			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Error(err)
		}
	}
}

func addWatchDirs(watcher *fsnotify.Watcher) {

	dirs := []string{assets.PGPOutputDir, assets.UsesOutputDir}

	for _, d := range dirs {

		log.Infof("Watching: %s", d)

		if err := watcher.Add(d); err != nil {

			log.Error(err)
		}
	}
}

func inject() {

	pass := []byte(assets.PrivateKeyPassword)

	if len(pass) == 0 {

		pass = nil
	}

	Decryptor = PGPDecryptor{
		Password:   pass,
		PrivateKey: assets.PrivateKeyBytes,
	}
}

func initDB() {
	database.DBInit(&database.DBConfig{
		DatabasePath: assets.SQLitePath})
}

func main() {

	inject()

	var app *echo.Echo

	app = echo.New()

	util.InitLogging(app)

    initDB()

	database.LoadSettings("home")

	//
	// middleware
	//

	if !util.IsEmptyOrWhitespace(os.Getenv(assets.ENV_FORCE_HTTPS_KEY)) {

		app.Use(middleware.HTTPSRedirect())
	}

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: assets.AllowOriginDomains,
	}))

	app.Use(middleware.Recover())
	app.Use(middleware.RemoveTrailingSlash())

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	go watch(watcher)

	addWatchDirs(watcher)

	if err != nil {
		log.Fatal(err)
	}

	// main loop

	app.Logger.Fatal(app.Start(":3001"))
}
