package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/assets"
	"github.com/minnowo/astoryofand/handler/crypto"
	"github.com/minnowo/astoryofand/model"
	"github.com/minnowo/astoryofand/view/order"
)

type OrderHandler struct {
}

func (h *OrderHandler) HandleOrderShow(c echo.Context) error {
	return render(c, order.Show(assets.BoxSetPrice, assets.StickerCost))
}

func (h *OrderHandler) HandleOrderThankYou(c echo.Context) error {
	return render(c, order.Show(assets.BoxSetPrice, assets.StickerCost))
}

func (h *OrderHandler) HandleOrderPlaced(c echo.Context) error {

	var o model.Order

	if err := c.Bind(&o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid order!")
	}

	if !o.CheckValidOrder() {
		return echo.NewHTTPError(http.StatusBadRequest, "This is an invalid order!")
	}

	o.TotalCost = float32(o.BoxSetCount)*assets.BoxSetPrice + float32(o.StickerCount)*assets.StickerCost

	log.Debug("Got order: ", o)

	jsonData, err := json.MarshalIndent(&o, "", "  ")

	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	}

	if err := crypto.WritePGPOrder(jsonData); err != nil {
		log.Debug(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error!")
	}

	return c.Redirect(http.StatusPermanentRedirect, "/order/thanks")
}
