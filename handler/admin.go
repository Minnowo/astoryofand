package handler

import (
	"crypto/subtle"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minnowo/astoryofand/database/memorydb"
	"github.com/minnowo/astoryofand/model"
	"github.com/minnowo/astoryofand/util"
	"github.com/minnowo/astoryofand/view/pages/admin"
)

type AdminHandler struct {
	Username []byte
	Password []byte
}

func (a *AdminHandler) HandleUserPasswordAdminAuth(username, password string, c echo.Context) (bool, error) {

	if subtle.ConstantTimeCompare([]byte(username), a.Username) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), a.Password) == 1 {
		return true, nil
	}

	return false, nil
}

func (a *AdminHandler) GetAdminPanel(c echo.Context) error {
	return util.EchoRenderTempl(c, admin.ShowAdminPane(&model.AdminView{
		BoxSetPrice: memorydb.GetDB().GetBoxPrice(),
		StickerCost: memorydb.GetDB().GetStickerPrice(),
	}))
}

func (a *AdminHandler) UpdateBoxPrice(c echo.Context) error {

	var o model.AdminView

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid request!")
	}

	memorydb.GetDB().SetBoxPrice(o.BoxSetPrice)

	return c.Redirect(http.StatusPermanentRedirect, "/admin")
}

func (a *AdminHandler) UpdateStickerPrice(c echo.Context) error {

	var o model.AdminView

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid request!")
	}

	memorydb.GetDB().SetStickerPrice(o.StickerCost)

	return c.Redirect(http.StatusPermanentRedirect, "/admin")
}
