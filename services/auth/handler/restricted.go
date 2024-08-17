package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Restricted(c echo.Context) error {
	return c.String(200, "You have access!")
}
