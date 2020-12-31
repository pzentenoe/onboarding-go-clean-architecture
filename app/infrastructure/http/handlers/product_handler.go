package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/pzentenoe/onboarding-go-clean-architecture/app/application/usecase"
	"net/http"
)

type productHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(e *echo.Echo, productUsecase usecase.ProductUsecase) *productHandler {
	h := &productHandler{productUsecase: productUsecase}

	g := e.Group("/products")
	g.GET("", h.getProducts)

	return h
}

func (h *productHandler) getProducts(c echo.Context) error {

	products, err := h.productUsecase.GetProducts()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, products)
}
