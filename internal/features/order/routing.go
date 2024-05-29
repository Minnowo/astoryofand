package order

import "github.com/labstack/echo/v4"

func (h *OrderHandler) Mount(e *echo.Echo) {

	o := e.Group("/order")

	o.GET("", h.HandleOrderShow)
	o.GET("/", h.HandleOrderShow)

	o.GET("/thanks", h.HandleOrderThankYou)

	// o.POST("/place", h.HandleOrderPlaced)
}
