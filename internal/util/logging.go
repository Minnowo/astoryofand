package util

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

const LOG_FMT = "${time_rfc3339_nano} ${remote_ip} ${level}"
const LOG_FMT_MIDDLEWARE = "${time_rfc3339} ${remote_ip} ${method} ${status} ${uri} \n"

func InitLogging(app *echo.Echo) {

	IS_DEBUG := os.Getenv("DEBUG")
	LOG_LEVEL := os.Getenv("LOG_LEVEL")

	// log.SetHeader(LOG_FMT)
	log.SetLevel(log.INFO)

	app.Logger.SetHeader(LOG_FMT)

	app.Use(middleware.Logger())
	// app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: LOG_FMT_MIDDLEWARE,
	// }))

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
