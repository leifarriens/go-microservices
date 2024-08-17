package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/leifarriens/go-microservices/services/product/model"
	_ "gorm.io/gorm"
)

type CreateProductParams struct {
	Name      string  `json:"name" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
	Available bool    `json:"available" validate:"required"`
}

// CreateProduct godoc
//
//	@Summary		Create product
//	@Description	Create product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//
//	@Param			product	body		model.ProductDto	true	"The input product struct"
//	@Success		200		{object}	model.ProductResponse
//
//	@failure		400		{string}	string	"error"
//	@failure		404		{string}	string	"error"
//	@failure		500		{string}	string	"error"
//
//	@Router			/products [post]
func (h *Handler) CreateProduct(c echo.Context) error {
	var p CreateProductParams

	if err := c.Bind(&p); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := c.Validate(p); err != nil {
		return err
	}

	id, err := h.ProductRepository.Create(c.Request().Context(), &model.Product{
		Name:      p.Name,
		Price:     p.Price,
		Available: p.Available,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	product, err := h.ProductRepository.FindById(c.Request().Context(), strconv.FormatUint(uint64(*id), 10))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if product == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, product)
}

// GetAllProducts godoc
//
//	@Summary		Get all products
//	@Description	Get all products
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.ProductResponse
//
//	@failure		400	{string}	string	"error"
//	@failure		404	{string}	string	"error"
//	@failure		500	{string}	string	"error"
//
//	@Router			/products [get]
func (h *Handler) GetAllProducts(c echo.Context) error {
	products, err := h.ProductRepository.FindAll(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}

type GetByIdParams struct {
	ID string `param:"id" validate:"required"`
}

// GetProduct godoc
//
//	@Summary		Get products by id
//	@Description	Get products by id
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Product ID"
//	@Success		200	{object}	model.ProductResponse
//
//	@failure		400	{string}	string	"error"
//	@failure		404	{string}	string	"error"
//	@failure		500	{string}	string	"error"
//
//	@Router			/products/{id} [get]
func (h *Handler) GetById(c echo.Context) error {
	var p GetByIdParams

	if err := c.Bind(&p); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := c.Validate(p); err != nil {
		return err
	}

	product, err := h.ProductRepository.FindById(c.Request().Context(), p.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if product == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, product)
}
