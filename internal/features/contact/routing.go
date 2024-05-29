package contact

import "github.com/labstack/echo/v4"

func (h *ContactUsHandler) Mount(e *echo.Echo) {

	o := e.Group("/contact")

	o.POST("/place", h.POSTContactUsPlaced)
}
