package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/assets"
	"github.com/minnowo/astoryofand/internal/crypto"
	"github.com/minnowo/astoryofand/internal/database"
	"github.com/minnowo/astoryofand/internal/database/models"
	"github.com/minnowo/astoryofand/internal/features/admin"
	"github.com/minnowo/astoryofand/internal/features/contact"
	"github.com/minnowo/astoryofand/internal/features/home"
	"github.com/minnowo/astoryofand/internal/features/order"
	"github.com/minnowo/astoryofand/internal/features/uses"
	"github.com/minnowo/astoryofand/internal/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func sanityCheck() {

	// sanity check, I don't want to accidentally include the private key in here
	if len(assets.PrivateKeyBytes) != 0 {
		panic("assets.PrivateKeyBytes has non-zero length! The Private Key may be embeded!")
	}

	if len(assets.PrivateKeyPassword) != 0 {
		panic("assets.PrivateKeyBytes has non-zero length! The Private Key may be embeded!")
	}
}

func initDB() {

	conf := &database.SqliteDBConf{
		DatabasePath: assets.SQLitePath,
	}
	database.DBInit(sqlite.Open(conf.GetDSN()), &gorm.Config{})
}

func main() {

	sanityCheck()

	var app *echo.Echo
	var orderEncryption crypto.EncryptionWriter
	var contactEncryption crypto.EncryptionWriter
	var usesEncryption crypto.EncryptionWriter

	app = echo.New()
	app.HideBanner = true
	app.HidePort = true

	util.InitLogging(app)

	initDB()

	database.LoadSettings("main")

	if username, ok := os.LookupEnv(assets.ENV_ADMIN_USERNAME_KEY); ok {

		if password, ok := os.LookupEnv(assets.ENV_ADMIN_PASSWORD_KEY); ok {

			log.Info("Creating user from env vars")

			if !database.InsertRawUser(&models.User{
				Username: username,
				Password: password,
			}) {
				log.Fatalf("Could not create user %s", username)
			}
		}
	}

	orderEncryption = &crypto.PGPEncryptionWriter{
		PublicKey:       assets.PublicKeyBytes,
		OutputDirectory: assets.PGPOutputDir,
	}

	contactEncryption = &crypto.PGPEncryptionWriter{
		PublicKey:       assets.PublicKeyBytes,
		OutputDirectory: assets.ContactOutputDir,
	}

	usesEncryption = &crypto.PGPEncryptionWriter{
		PublicKey:       assets.PublicKeyBytes,
		OutputDirectory: assets.UsesOutputDir,
	}

	orderEncryption.EnsureCanWriteDiskOrExit()
	contactEncryption.EnsureCanWriteDiskOrExit()
	usesEncryption.EnsureCanWriteDiskOrExit()

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

	//
	// admin routes
	//

	adminHandler := admin.AdminHandler{}
	adminHandler.Mount(app)

	//
	// order routes
	//

	orderHandler := order.OrderHandler{
		EncryptionWriter: orderEncryption,
	}
	orderHandler.Mount(app)

	//
	// contact us routes
	//
	contactusHandler := contact.ContactUsHandler{
		EncryptionWriter: contactEncryption,
	}
	contactusHandler.Mount(app)

	//
	// usecases routes
	//

	usesHandler := uses.UsesHandler{
		EncryptionWriter: usesEncryption,
	}
	usesHandler.Mount(app)

	//
	// home routes
	//

	commonHandler := home.HomeHandler{}
	commonHandler.Mount(app)

	// main loop

	port := ":3000"
	log.Info("Running server on port", port)

	app.Logger.Fatal(app.Start(port))
}
