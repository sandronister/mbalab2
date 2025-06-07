package web

import (
	"github.com/labstack/echo"
	"github.com/sandronister/mbalab2/weather_service/internal/infra/handler"
)

func (s *Server) RegisterRoutes(h *handler.WeatherByCep) {
	s.router.GET("/weather/:cep", h.Get)
	s.router.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	})
}
