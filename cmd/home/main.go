package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/assets"
	"github.com/minnowo/astoryofand/internal/database"
	"github.com/minnowo/astoryofand/internal/database/models"
	"github.com/minnowo/astoryofand/internal/templates/pages"
	"github.com/minnowo/astoryofand/internal/util"
)

type DecrypFailRead int32

const (
	SUCESS DecrypFailRead = 1 << iota
	NOT_THE_TYPE_OF_FILE
	COULD_NOT_READ_FILE
	SOME_ERROR_WHILE_DECRYPTING
	JSON_MARSHAL
)

type FileToDecrypt struct {
	Path       string
	Tried      int32
	AbortAfter int32
}

type PGPDecryptor struct {
	PrivateKey string
	Password   []byte
	FileQ      chan (string)
	Files      map[string]*FileToDecrypt
}

var (
	Decryptor PGPDecryptor
)

func (pgp *PGPDecryptor) run() {

	log.Info("Decryptor is running...")

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {

		select {

		case file, ok := <-pgp.FileQ:

			if !ok {

				return
			}

			pgp.Files[file] = &FileToDecrypt{
				Path:       file,
				Tried:      0,
				AbortAfter: 5,
			}

			break

		case <-ticker.C:

			log.Info("Decryptor Tick")

			keys := make([]string, len(pgp.Files))

			i := 0
			for k := range pgp.Files {
				keys[i] = k
				i++
			}

			for _, path := range keys {

				file := pgp.Files[path]

				switch pgp.processFile(path) {

				case NOT_THE_TYPE_OF_FILE:
					delete(pgp.Files, path)
					break

				case SUCESS:
					log.Infof("File %s was added to the database", path)
					delete(pgp.Files, path)
					break

				case JSON_MARSHAL:
					log.Errorf("File %s has a JSON Marhsal problem", path)
					delete(pgp.Files, path)
					break

				case SOME_ERROR_WHILE_DECRYPTING:
				case COULD_NOT_READ_FILE:

					file.Tried++

					if file.Tried > file.AbortAfter {
						log.Infof("Removing %s because it failed %d times", path, file.Tried)
						delete(pgp.Files, path)
					}
				}
			}

			break

		}
	}
}

func (pgp *PGPDecryptor) processFile(path string) DecrypFailRead {

	if filepath.Ext(path) != ".asc" {

		log.Warnf("Will not process %s because it does not end with .asc", path)

		return NOT_THE_TYPE_OF_FILE
	}

	log.Debug("Opening file...")

	fileContent, err := os.ReadFile(path)

	if err != nil {

		log.Error(err)

		return COULD_NOT_READ_FILE
	}

	log.Debug("Decrypting file...")

	data, err := helper.DecryptBinaryMessageArmored(pgp.PrivateKey, pgp.Password, string(fileContent))

	if err != nil {

		log.Errorf("Could not decrpt file: %v", err)

		return SOME_ERROR_WHILE_DECRYPTING
	}

	log.Debug(data)

	log.Debug("Unmarshal file...")

	var e models.UserData

	if err := json.Unmarshal(data, &e); err != nil {

		log.Error(err)

		return JSON_MARSHAL
	}

	log.Debug(e)

	switch e.Type {

	case models.OrderType:

		log.Info("Got Order from file")

		var o models.Order

		if err := json.Unmarshal(data, &o); err != nil {

			log.Error(err)

			return JSON_MARSHAL
		}

		log.Debug(o)

		database.InsertOrder(&o)

		break

	case models.UsecaseType:

		log.Info("Got Usecase from file")

		var o models.UseCase

		if err := json.Unmarshal(data, &o); err != nil {

			log.Error(err)

			return JSON_MARSHAL
		}

		log.Debug(o)
		database.InsertUseCase(&o)

		break
	}

	return SUCESS
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

				Decryptor.FileQ <- event.Name

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
		FileQ:      make(chan (string), 32),
		Files:      make(map[string]*FileToDecrypt),
	}
	go Decryptor.run()
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

	//
	// static assets
	//

	staticAssetHandler := http.FileServer(assets.GetFileSystem(false))

	app.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", staticAssetHandler)))
	app.GET("/robots.txt", echo.WrapHandler(staticAssetHandler))
	app.GET("/favicon.ico", echo.WrapHandler(staticAssetHandler))
	app.GET("/favicon.png", echo.WrapHandler(staticAssetHandler))

	app.GET("", func(c echo.Context) error {
		return util.EchoRenderTempl(c, pages.ShowInvoice(models.NewOrder()))
	})

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
