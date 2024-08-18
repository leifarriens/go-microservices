package handler

import (
	"crypto/rsa"

	"github.com/labstack/echo/v4"
	"github.com/leifarriens/go-microservices/internal/shared"
	"github.com/leifarriens/go-microservices/services/product/service"
)

type Handler struct {
	ProductService service.ProductService
	PublicKey      *rsa.PublicKey
}

type HandlerConfig struct {
	E              *echo.Echo
	ProductService service.ProductService
	PublicKey      *rsa.PublicKey
}

func NewHandler(config *HandlerConfig) *Handler {
	h := &Handler{
		ProductService: config.ProductService,
		PublicKey:      config.PublicKey,
	}

	config.E.POST("/products", h.CreateProduct, shared.Authorize(h.PublicKey))
	config.E.GET("/products", h.GetAllProducts)
	config.E.GET("/products/:id", h.GetById)

	return h
}
