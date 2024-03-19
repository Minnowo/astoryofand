package admin

import (
	"crypto/subtle"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minnowo/astoryofand/internal/database"
	"github.com/minnowo/astoryofand/internal/database/models"
	"github.com/minnowo/astoryofand/internal/templates/pages/admin"
	"github.com/minnowo/astoryofand/internal/util"
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
	return util.EchoRenderTempl(c, admin.ShowAdminPane(&models.AdminView{
		BoxSetPrice: database.GetBoxPrice(),
		StickerCost: database.GetStickerPrice(),
	}))
}

func (a *AdminHandler) UpdateBoxPrice(c echo.Context) error {

	var o models.AdminView

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid request!")
	}

	database.SetBoxPrice(o.BoxSetPrice)

	return c.Redirect(http.StatusPermanentRedirect, "/admin")
}

func (a *AdminHandler) UpdateStickerPrice(c echo.Context) error {

	var o models.AdminView

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid request!")
	}

	database.SetStickerPrice(o.StickerCost)

	return c.Redirect(http.StatusPermanentRedirect, "/admin")
}
