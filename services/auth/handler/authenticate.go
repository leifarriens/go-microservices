package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Authenticate godoc
//
//	@Tags		auth
//	@Success	200
//	@failure	500	{string}	string	"error"
//	@Router		/authenticate [post]
func (h *Handler) Authenticate(c echo.Context) error {
	accessToken, err := h.TokenService.GenerateAccessToken(c.Request().Context())

	if err != nil {
		return echo.ErrInternalServerError
	}

	accessTokenCookie := &http.Cookie{
		Name:     "accessToken",
		Value:    accessToken.Token,
		Expires:  accessToken.Expires,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		Domain:   h.Domain,
	}

	c.SetCookie(accessTokenCookie)

	return c.JSON(http.StatusOK, echo.Map{
		"accessToken": accessToken.Token,
	})
}
