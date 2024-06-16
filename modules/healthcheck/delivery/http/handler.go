package healthcheck

import (
	"base-golang-echo-setup/modules/healthcheck/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	UseCase healthcheck.UseCase
}

// NewHandler will initialize the health check endpoint
func NewHandler(e *echo.Echo, u healthcheck.UseCase) {
	handler := &Handler{
		UseCase: u,
	}
	e.GET("/v1/health-check", handler.CheckHealth)
}

// CheckHealth will fetch the article based on given params
func (h *Handler) CheckHealth(c echo.Context) error {
	res, err := h.UseCase.CheckHealth()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, res)
}
