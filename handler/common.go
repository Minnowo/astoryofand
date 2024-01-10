package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/minnowo/astoryofand/util"
	"github.com/minnowo/astoryofand/view/pages"
)

type CommonHandler struct {
}

func (h *CommonHandler) HandleLicenseShow(c echo.Context) error {
	return util.EchoRenderTempl(c, pages.ShowLicensePage())
}

func (h *CommonHandler) HandleAboutShow(c echo.Context) error {
	return util.EchoRenderTempl(c, pages.ShowAboutPage())
}
