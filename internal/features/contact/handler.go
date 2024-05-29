package contact

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/database/models"
	"github.com/minnowo/astoryofand/internal/util"
)

func (h *ContactUsHandler) POSTContactUsPlaced(c echo.Context) error {

	var o models.TableContact

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid contact!")
	}

	if err := o.CheckValidDataFromUser(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	o.DelayedInit()

	log.Debug("Someone is getting in contact")

	jsonData, err := json.MarshalIndent(o.EnsureType(), "", "  ")

	if err != nil {

		log.Error(err)

		go util.SendDiscordOrderWebhook("ERROR Marshaling contact")

		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	}

	if oid, err := h.EncryptionWriter.SaveAndEncryptData(o.UUID, jsonData); err != nil {

		log.Error(err)

		go util.SendDiscordOrderWebhook("ERROR Encrypting contact")

		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	} else {

		go util.SendDiscordOrderWebhook("Someone is trying to get in contact. id: `" + oid + "`")

		return c.Redirect(http.StatusSeeOther, "/order")
	}

}
