package main

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/minnowo/astoryofand/handler"
	"github.com/minnowo/astoryofand/handler/crypto"
	"github.com/minnowo/astoryofand/model"
)

func initLogging(app *echo.Echo) {

	IS_DEBUG := os.Getenv("DEBUG")
	LOG_LEVEL := os.Getenv("LOG_LEVEL")

	log.SetHeader("${time_rfc3339} ${level}")

	if l, ok := app.Logger.(*log.Logger); ok {

		l.SetHeader("${time_rfc3339} ${level}")

		if !model.IsEmptyOrWhitespace(LOG_LEVEL) {

			if level, err := strconv.ParseUint(LOG_LEVEL, 10, 8); err == nil {
				l.SetLevel(log.Lvl(level))
				log.SetLevel(log.Lvl(level))
				l.Info("Read LOG_LEVEL from env: ", level)
			} else {
				l.Error("Could not read LOG_LEVEL from env. Log level is: ", l.Level())
			}
		}
	}

	if debug_, err := strconv.ParseBool(IS_DEBUG); err == nil {
		app.Debug = debug_
		app.Logger.Info("Read DEBUG from Env: ", debug_)
	} else {
		app.Debug = false
		app.Logger.Error("Could not read DEBUG from env. Running in release mode.")
	}
}

func main() {
	crypto.FailIfPGPDirNotExists()

	app := echo.New()

	initLogging(app)

	// app.Use(middleware.HTTPSRedirect())
	app.Use(middleware.Recover())

	app.Static("/static", "static")

	orderHandler := handler.OrderHandler{}
	app.Any("/order", orderHandler.HandleOrderShow)
	app.Any("/order/thanks", orderHandler.HandleOrderThankYou)
	app.POST("/order/place", orderHandler.HandleOrderPlaced)

	commonHandler := handler.CommonHandler{}
	app.Any("/license", commonHandler.HandleLicenseShow)
	app.Any("/about", commonHandler.HandleAboutShow)

	app.RouteNotFound("/*", orderHandler.HandleOrderShow)

	app.Logger.Fatal(app.Start(":3000"))
}
