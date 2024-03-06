package uses

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/model"
	"github.com/minnowo/astoryofand/internal/templates/pages/uses"
	"github.com/minnowo/astoryofand/internal/util"
)

func (u *UsesHandler) HandleUsesGET(c echo.Context) error {
	return util.EchoRenderTempl(c, uses.ShowUsesPage())
}

func (u *UsesHandler) HandleUsesThankYouGET(c echo.Context) error {

	return util.EchoRenderTempl(c, uses.ShowUsesThanks())
}

func (u *UsesHandler) HandleUsesPOST(c echo.Context) error {

	var o model.UseCase

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid use!")
	}

	if !o.CheckValid() {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid use!")
	}

	log.Debug("Got Use: ", o)

	jsonData, err := json.MarshalIndent(&o, "", "  ")

	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	}

	if _, err := u.EncryptionWriter.SaveAndEncryptData(jsonData); err != nil {

		log.Debug(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	} else {

		go util.SendDiscordOrderWebhook("Someone has given a Use Case!")

		return c.Redirect(http.StatusPermanentRedirect, "/uses/thanks")
	}

}
