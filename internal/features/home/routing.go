package home

import "github.com/labstack/echo/v4"

func (h *HomeHandler) Mount(e *echo.Echo) {

	e.GET("", h.HandleHome)
	e.GET("/", h.HandleHome)

	e.GET("/home", h.HandleHome)

	e.GET("/license", h.HandleLicenseShow)

	e.GET("/about", h.HandleAboutShow)

	e.GET("/grid_test_with_audio", h.HandleGride)
}
