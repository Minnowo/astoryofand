package main

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/minnowo/astoryofand/assets"
	"github.com/minnowo/astoryofand/database/memorydb"
	"github.com/minnowo/astoryofand/handler"
	"github.com/minnowo/astoryofand/handler/crypto"
	"github.com/minnowo/astoryofand/util"
)

func initLogging(app *echo.Echo) {

	IS_DEBUG := os.Getenv("DEBUG")
	LOG_LEVEL := os.Getenv("LOG_LEVEL")

	log.SetHeader("${time_rfc3339} ${level}")
	log.SetLevel(log.INFO)

	app.Logger.SetHeader("${time_rfc3339} ${level}")

	if level, err := strconv.ParseUint(LOG_LEVEL, 10, 8); err == nil {
		app.Logger.SetLevel(log.Lvl(level))
		log.Info("Read LOG_LEVEL from env: ", level)
	} else {
		log.Warn("Could not read LOG_LEVEL from env. Log level is: ", app.Logger.Level())
	}

	if debug_, err := strconv.ParseBool(IS_DEBUG); err == nil {

		app.Debug = debug_

		if debug_ {
			app.Logger.SetLevel(log.DEBUG)
		}

		log.Info("Read DEBUG from Env: ", debug_)
	} else {
		app.Debug = false
		log.Warn("Could not read DEBUG from env. Running in release mode.")
	}

	log.SetLevel(app.Logger.Level())
}

func main() {

	// sanity check, I don't want to accidentally include the private key in here
	if len(assets.PrivateKeyBytes) != 0 {
		panic("assets.PrivateKeyBytes has non-zero length! The Private Key may be embeded!")
	}

	var app *echo.Echo
	var orderEncryption crypto.EncryptionWriter
	var usesEncryption crypto.EncryptionWriter

	app = echo.New()

	initLogging(app)

	memorydb.InitDB()

	orderEncryption = &crypto.PGPEncryptionWriter{
		PublicKey:       assets.PublicKeyBytes,
		OutputDirectory: assets.PGPOutputDir,
	}

	usesEncryption = &crypto.PGPEncryptionWriter{
		PublicKey:       assets.PublicKeyBytes,
		OutputDirectory: assets.UsesOutputDir,
	}

	orderEncryption.EnsureCanWriteDiskOrExit()
	usesEncryption.EnsureCanWriteDiskOrExit()

	if !util.IsEmptyOrWhitespace(os.Getenv(assets.ENV_FORCE_HTTPS_KEY)) {
		app.Use(middleware.HTTPSRedirect())
	}
	app.Use(middleware.Recover())

	app.Static("/static", "static")
	app.File("/robots.txt", "static/robots.txt")

	adminHandler := handler.AdminHandler{
		Username: []byte(os.Getenv(assets.ENV_ADMIN_USERNAME_KEY)),
		Password: []byte(os.Getenv(assets.ENV_ADMIN_PASSWORD_KEY)),
	}
	admin := app.Group("/admin")
	admin.Use(middleware.BasicAuth(adminHandler.HandleUserPasswordAdminAuth))
	admin.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(3)))
	admin.Use(middleware.Logger())
	admin.Any("", adminHandler.GetAdminPanel)
	admin.Any("/", adminHandler.GetAdminPanel)
	admin.POST("/update/boxprice", adminHandler.UpdateBoxPrice)
	admin.POST("/update/stickerprice", adminHandler.UpdateStickerPrice)

	orderHandler := handler.OrderHandler{
		EncryptionWriter: orderEncryption,
	}
	order := app.Group("/order")
	order.Any("", orderHandler.HandleOrderShow)
	order.Any("/", orderHandler.HandleOrderShow)
	order.Any("/thanks", orderHandler.HandleOrderThankYou)
	order.POST("/place", orderHandler.HandleOrderPlaced)

	usesHandler := handler.UsesHandler{
		EncryptionWriter: usesEncryption,
	}
	uses := app.Group("/uses")
	uses.Any("", usesHandler.HandleUsesGET)
	uses.Any("/", usesHandler.HandleUsesGET)
	uses.Any("/thanks", usesHandler.HandleUsesThankYouGET)
	uses.POST("/place", usesHandler.HandleUsesPOST)

	commonHandler := handler.CommonHandler{}
	app.Any("", commonHandler.HandleHome)
	app.Any("/", commonHandler.HandleHome)
	app.Any("/home", commonHandler.HandleHome)
	app.Any("/license", commonHandler.HandleLicenseShow)
	app.Any("/about", commonHandler.HandleAboutShow)

	app.Logger.Fatal(app.Start(":3000"))
}
