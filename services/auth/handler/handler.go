package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/leifarriens/go-microservices/internal/shared"
	"github.com/leifarriens/go-microservices/services/auth/service"
)

type Handler struct {
	TokenService service.TokenService
	Domain       string
}

type HandlerConfig struct {
	E            *echo.Echo
	TokenService service.TokenService
	Domain       string
}

func NewHandler(config *HandlerConfig) *Handler {
	h := &Handler{
		TokenService: config.TokenService,
		Domain:       config.Domain,
	}

	config.E.POST("/authenticate", h.Authenticate)
	config.E.POST("/logout", h.Logout)
	config.E.GET("/restricted", h.Restricted, shared.Authorize(h.TokenService.GetPublicKey()))

	return h
}
