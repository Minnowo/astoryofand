package uses

import (
	"github.com/labstack/echo/v4"
)

func (h *UsesHandler) Mount(e *echo.Echo) {

	u := e.Group("/uses")

	u.GET("", h.HandleUsesGET)
	u.GET("/", h.HandleUsesGET)

	u.GET("/thanks", h.HandleUsesThankYouGET)

	u.POST("/place", h.HandleUsesPOST)
}
