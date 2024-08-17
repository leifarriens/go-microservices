package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/leifarriens/go-microservices/services/product/repository"
)

type Handler struct {
	ProductService repository.ProductRepository
}

type HandlerConfig struct {
	E                 *echo.Echo
	ProductRepository repository.ProductRepository
}

func NewHandler(config *HandlerConfig) *Handler {
	h := &Handler{
		ProductService: config.ProductRepository,
	}

	config.E.POST("/products", h.CreateProduct)
	config.E.GET("/products", h.GetAllProducts)
	config.E.GET("/products/:id", h.GetById)

	return h
}
