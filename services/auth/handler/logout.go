package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Logout godoc
//
//	@Tags		auth
//	@Success	200
//	@failure	500	{string}	string	"error"
//	@Router		/logout [post]
func (h *Handler) Logout(c echo.Context) error {
	accessTokenCookie := &http.Cookie{
		Name:    "accessToken",
		Value:   "",
		Expires: time.Now().Add(-1),
		Domain:  h.Domain,
	}

	c.SetCookie(accessTokenCookie)

	return c.NoContent(200)
}
