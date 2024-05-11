package admin

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *AdminHandler) Mount(e *echo.Echo) {

	g := e.Group("/admin")

	g.Use(middleware.BasicAuth(h.HandleUserPasswordAdminAuth))
	g.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(3)))
	g.Use(middleware.Logger())

	g.Match([]string{"GET", "POST"}, "", h.GetAdminPanel)

	g.Match([]string{"GET", "POST"}, "/", h.GetAdminPanel)

	g.POST("/update/boxprice", h.UpdateBoxPrice)

	g.POST("/update/stickerprice", h.UpdateStickerPrice)

	g.POST("/create/user", h.POSTCreateUser)
}
