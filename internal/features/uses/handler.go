package uses

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/database/models"
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

	var o models.TableUseCase

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid use!")
	}

	if err := o.CheckValidDataFromUser(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	o.DelayedInit()

	log.Debug("Got Use: ", o.EnsureType())

	jsonData, err := json.MarshalIndent(o.EnsureType(), "", "  ")

	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	}

	if _, err := u.EncryptionWriter.SaveAndEncryptData(o.UUID, jsonData); err != nil {

		log.Debug(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	} else {

		go util.SendDiscordOrderWebhook("Someone has given a Use Case!")

		return c.Redirect(http.StatusSeeOther, "/uses/thanks")
	}

}
