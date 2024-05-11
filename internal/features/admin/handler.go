package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minnowo/astoryofand/internal/database"
	"github.com/minnowo/astoryofand/internal/database/models"
	"github.com/minnowo/astoryofand/internal/templates/pages/admin"
	"github.com/minnowo/astoryofand/internal/util"
)

type AdminHandler struct {
}

func (a *AdminHandler) HandleUserPasswordAdminAuth(username, password string, c echo.Context) (bool, error) {

	var usr models.User

	usr.Username = username
	usr.Password = password

	return database.AuthUser(&usr)
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

	return a.GetAdminPanel(c)
}

func (a *AdminHandler) UpdateStickerPrice(c echo.Context) error {

	var o models.AdminView

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid request!")
	}

	database.SetStickerPrice(o.StickerCost)

	return a.GetAdminPanel(c)
}

func (a *AdminHandler) POSTCreateUser(c echo.Context) error {

	var o models.User

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid request!")
	}

	if !o.CheckValid() {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user details")
	}

	database.InsertRawUser(&o)

	return a.GetAdminPanel(c)
}
