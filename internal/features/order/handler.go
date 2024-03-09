package order

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/internal/database"
	"github.com/minnowo/astoryofand/internal/model"
	"github.com/minnowo/astoryofand/internal/templates/pages/order"
	"github.com/minnowo/astoryofand/internal/util"
)

func (h *OrderHandler) HandleOrderShow(c echo.Context) error {
	return util.EchoRenderTempl(c, order.ShowOrderPage(database.GetBoxPrice(), database.GetStickerPrice()))
}

func (h *OrderHandler) HandleOrderThankYou(c echo.Context) error {
	return util.EchoRenderTempl(c, order.ShowOrderThanks(c.QueryParam("oid")))
}

func (h *OrderHandler) HandleOrderPlaced(c echo.Context) error {

	var o model.Order

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid order!")
	}

	if err := o.CheckValid(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// we are safe to use the o.StickerValue and o.BoxSetValue here because the above
	// valid check ensures they are equal to the value stored in the database
	o.TotalCost = float32(o.BoxSetCount)*o.BoxSetValue + float32(o.StickerCount)*o.StickerValue

	log.Debug("Got order: ", o)

	jsonData, err := json.MarshalIndent(&o, "", "  ")

	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	}

	if oid, err := h.EncryptionWriter.SaveAndEncryptData(jsonData); err != nil {

		log.Debug(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	} else {

		go util.SendDiscordOrderWebhook("New order with id: `" + oid + "`")

		params := url.Values{}

		params.Add("oid", oid)

		return c.Redirect(http.StatusPermanentRedirect, "/order/thanks?"+params.Encode())
	}

}
