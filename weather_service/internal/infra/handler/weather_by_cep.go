package handler

import (
	"github.com/labstack/echo"
	"github.com/sandronister/mbalab2/weather_service/internal/usecase/types"
)

type WeatherByCep struct {
	cepUsecase     types.ICep
	weatherUsecase types.IWeatherUseCase
}

func NewWeatherByCep(cepUsecase types.ICep, weatherUseCase types.IWeatherUseCase) *WeatherByCep {
	return &WeatherByCep{
		cepUsecase:     cepUsecase,
		weatherUsecase: weatherUseCase,
	}
}

func (w *WeatherByCep) Get(c echo.Context) error {
	cep := c.Param("cep")

	resp, err := w.cepUsecase.Get(cep)
	if err != nil {
		return c.JSON(err.StatusCode, map[string]string{
			"error": err.Error(),
		})
	}

	weatherResult, err := w.weatherUsecase.Get(resp.Localidade)

	if err != nil {
		return c.JSON(err.StatusCode, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"data": weatherResult,
	})
}
