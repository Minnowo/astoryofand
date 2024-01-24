package main

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/minnowo/astoryofand/assets"
	"github.com/minnowo/astoryofand/handler"
	"github.com/minnowo/astoryofand/handler/crypto"
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

	var app *echo.Echo
	var orderEncryption crypto.EncryptionWriter
	var usesEncryption crypto.EncryptionWriter

	app = echo.New()

	initLogging(app)

	orderEncryption = &crypto.PGPEncryptionWriter{
		PublicKey:       string(assets.PublicKeyBytes),
		OutputDirectory: assets.PGPOutputDir,
	}

	usesEncryption = &crypto.PGPEncryptionWriter{
		PublicKey:       string(assets.PublicKeyBytes),
		OutputDirectory: assets.UsesOutputDir,
	}

	orderEncryption.EnsureCanWriteDiskOrExit()
	usesEncryption.EnsureCanWriteDiskOrExit()

	// app.Use(middleware.HTTPSRedirect())
	app.Use(middleware.Recover())

	app.Static("/static", "static")

	orderHandler := handler.OrderHandler{
		EncryptionWriter: orderEncryption,
	}
	order := app.Group("/order")
	order.Any("", orderHandler.HandleOrderShow)
	order.Any("/thanks", orderHandler.HandleOrderThankYou)
	order.POST("/place", orderHandler.HandleOrderPlaced)

	usesHandler := handler.UsesHandler{
		EncryptionWriter: usesEncryption,
	}
	uses := app.Group("/uses")
	uses.Any("", usesHandler.HandleUsesGET)
	uses.Any("/thanks", usesHandler.HandleUsesThankYouGET)
	uses.POST("/place", usesHandler.HandleUsesPOST)

	commonHandler := handler.CommonHandler{}
	app.Any("/home", commonHandler.HandleHome)
	app.Any("/license", commonHandler.HandleLicenseShow)
	app.Any("/about", commonHandler.HandleAboutShow)

	app.RouteNotFound("/*", commonHandler.HandleHome)

	app.Logger.Fatal(app.Start(":3000"))
}
