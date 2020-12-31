package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type healthHandler struct {
}

func NewHealthHandler(e *echo.Echo) *healthHandler {
	h := new(healthHandler)

	e.GET("/health", h.healthCheck)
	return h
}

func (h *healthHandler) healthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"status": "UP", })
}
