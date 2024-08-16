package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllProducts(c echo.Context) error {
	products, err := h.ProductService.FindAll(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}
