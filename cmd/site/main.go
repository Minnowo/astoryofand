package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minnowo/astoryofand/internal/assets"
	"github.com/minnowo/astoryofand/internal/crypto"
	"github.com/minnowo/astoryofand/internal/database"
	"github.com/minnowo/astoryofand/internal/features/admin"
	"github.com/minnowo/astoryofand/internal/features/home"
	"github.com/minnowo/astoryofand/internal/features/order"
	"github.com/minnowo/astoryofand/internal/features/uses"
	"github.com/minnowo/astoryofand/internal/util"
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
	database.DBInit(&database.DBConfig{
		DatabasePath: assets.SQLitePath})
}

func main() {

	sanityCheck()

	var app *echo.Echo
	var orderEncryption crypto.EncryptionWriter
	var usesEncryption crypto.EncryptionWriter

	app = echo.New()

	util.InitLogging(app)

	initDB()

	database.LoadSettings("main")

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

	adminHandler := admin.AdminHandler{
		Username: []byte(os.Getenv(assets.ENV_ADMIN_USERNAME_KEY)),
		Password: []byte(os.Getenv(assets.ENV_ADMIN_PASSWORD_KEY)),
	}
	adminHandler.Mount(app)

	//
	// order routes
	//

	orderHandler := order.OrderHandler{
		EncryptionWriter: orderEncryption,
	}
	orderHandler.Mount(app)

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

	app.Logger.Fatal(app.Start(":3000"))
}
