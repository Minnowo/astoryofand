package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/handler/crypto"
	"github.com/minnowo/astoryofand/model"
	"github.com/minnowo/astoryofand/util"
	"github.com/minnowo/astoryofand/view/pages/uses"
)

type UsesHandler struct {
	EncryptionWriter crypto.EncryptionWriter
}

func (u *UsesHandler) HandleUsesGET(c echo.Context) error {
	return util.EchoRenderTempl(c, uses.ShowUsesPage())
}

func (u *UsesHandler) HandleUsesThankYouGET(c echo.Context) error {

	return util.EchoRenderTempl(c, uses.ShowUsesThanks())
}

func (u *UsesHandler) HandleUsesPOST(c echo.Context) error {

	var o model.Use

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

	if oid, err := u.EncryptionWriter.SaveAndEncryptData(jsonData); err != nil {

		log.Debug(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	} else {

		params := url.Values{}

		params.Add("oid", oid)

		return c.Redirect(http.StatusPermanentRedirect, "/uses/thanks?"+params.Encode())
	}

}
