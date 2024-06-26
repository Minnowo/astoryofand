package home

import (
	"github.com/labstack/echo/v4"
	"github.com/minnowo/astoryofand/internal/templates/pages"
	"github.com/minnowo/astoryofand/internal/util"
)

func (h *HomeHandler) HandleLicenseShow(c echo.Context) error {
	return util.EchoRenderTempl(c, pages.ShowLicensePage())
}

func (h *HomeHandler) HandleAboutShow(c echo.Context) error {
	return util.EchoRenderTempl(c, pages.ShowAboutPage())
}

func (h *HomeHandler) HandleHome(c echo.Context) error {
	return util.EchoRenderTempl(c, pages.ShowHomePage())
}

func (h *HomeHandler) HandleGride(c echo.Context) error {

	flat_view := c.QueryParam("flat")

	return util.EchoRenderTempl(c, pages.ShowGridPage(flat_view == "1"))
}
