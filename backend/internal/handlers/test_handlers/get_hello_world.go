package testhandlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetHelloWorld(c echo.Context) error {
	res, isFound := h.service.GetHelloWorld(c.Request().Context())
	if !isFound {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, res)
}
